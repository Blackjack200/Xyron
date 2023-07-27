<?php

namespace prokits\xyron\loader;

use Grpc\ChannelCredentials;
use libasync\await\Await;
use libasync\executor\Executor;
use libasync\executor\ThreadFactory;
use libasync\executor\ThreadPoolExecutor;
use libasync\utils\LoggerUtils;
use pocketmine\block\Air;
use pocketmine\event\block\BlockBreakEvent;
use pocketmine\event\block\BlockPlaceEvent;
use pocketmine\event\entity\EntityDamageByEntityEvent;
use pocketmine\event\entity\EntityEffectAddEvent;
use pocketmine\event\entity\EntityEffectRemoveEvent;
use pocketmine\event\entity\EntityMotionEvent;
use pocketmine\event\entity\EntityTeleportEvent;
use pocketmine\event\inventory\InventoryCloseEvent;
use pocketmine\event\inventory\InventoryOpenEvent;
use pocketmine\event\Listener;
use pocketmine\event\player\PlayerDeathEvent;
use pocketmine\event\player\PlayerGameModeChangeEvent;
use pocketmine\event\player\PlayerItemConsumeEvent;
use pocketmine\event\player\PlayerJoinEvent;
use pocketmine\event\player\PlayerJumpEvent;
use pocketmine\event\player\PlayerMoveEvent;
use pocketmine\event\player\PlayerQuitEvent;
use pocketmine\event\player\PlayerRespawnEvent;
use pocketmine\event\player\PlayerToggleFlightEvent;
use pocketmine\event\player\PlayerToggleGlideEvent;
use pocketmine\event\player\PlayerToggleSneakEvent;
use pocketmine\event\player\PlayerToggleSprintEvent;
use pocketmine\event\player\PlayerToggleSwimEvent;
use pocketmine\item\ItemTypeIds;
use pocketmine\math\AxisAlignedBB;
use pocketmine\math\Vector3;
use pocketmine\player\Player;
use pocketmine\plugin\Plugin;
use pocketmine\plugin\PluginBase;
use pocketmine\scheduler\ClosureTask;
use pocketmine\Server;
use pocketmine\world\World;
use prokits\xyron\AddPlayerRequest;
use prokits\xyron\AnticheatClient;
use prokits\xyron\AttackData;
use prokits\xyron\ConsumeStatus;
use prokits\xyron\EntityPositionData;
use prokits\xyron\Judgement;
use prokits\xyron\JudgementData;
use prokits\xyron\Player as XyronPlayer;
use prokits\xyron\PlayerAction;
use prokits\xyron\PlayerActionData;
use prokits\xyron\PlayerAttackData;
use prokits\xyron\PlayerBreakBlockData;
use prokits\xyron\PlayerEatFoodData;
use prokits\xyron\PlayerEffectData;
use prokits\xyron\PlayerGameModeData;
use prokits\xyron\PlayerInputModeData;
use prokits\xyron\PlayerLifeData;
use prokits\xyron\PlayerMotionData;
use prokits\xyron\PlayerMoveData;
use prokits\xyron\PlayerPlaceBlockData;
use prokits\xyron\PlayerReceipt;
use prokits\xyron\PlayerReport;
use prokits\xyron\ReportResponse;
use prokits\xyron\TimestampedReportData;
use RuntimeException;
use Symfony\Component\Filesystem\Path;
use const Grpc\STATUS_OK;

class Loader extends PluginBase implements Listener {
	private ThreadPoolExecutor $executor;
	private string $autoload;
	/** @var XyronData[] */
	protected array $data = [];

	public function __construct(...$args) {
		parent::__construct(...$args);
		$this->autoload = Path::join(dirname(__DIR__, 4), 'vendor', 'autoload.php');
		require_once $this->autoload;
	}

	protected function onEnable() : void {
		$this->getServer()->getPluginManager()->registerEvents($this, $this);
		$this->executor = self::createThreadPoolExecutor($this, $this->autoload, "localhost:8884");
		$this->executor->start();
		//more than 100000 years available LOL so PHP_INT_MAX
		Await::do(function() {
			Await::tick(function() : void {
				foreach (Server::getInstance()->getOnlinePlayers() as $player) {
					if (isset($this->data[spl_object_id($player)])) {
						$data = $this->data[spl_object_id($player)];
						$tData = [];
						foreach ($data->getQueue()->flush($this->getTick()) as $tck => $wdata) {
							$tData[$tck] = (new TimestampedReportData())->setData($wdata);
						}

						$rpData = (new PlayerReport())
							->setPlayer($data->receipt)
							->setLatency($player->getNetworkSession()->getPing() / 1000)
							->setData($tData)
							->serializeToString();

						$jdjm = new ReportResponse();
						$jdjm->mergeFromString(Await::fiberAsync(static function(AnticheatClient $client) use ($rpData) : string {
							$rp = (new PlayerReport());
							$rp->mergeFromString($rpData);
							[$resp, $status] = $client->Report($rp)->wait();
							if ($status->code !== STATUS_OK) {
								throw new RuntimeException($status);
							}
							assert($resp instanceof ReportResponse);
							return $resp->serializeToString();
						}, $this->executor));
						$this->handleJudgements($player, iterator_to_array($jdjm->getJudgements()->getIterator()));
					}
				}
			}, 4, PHP_INT_MAX);
		})->panic();
	}

	protected function onDisable() : void {
		$this->executor->shutdown();
	}

	public static function createThreadPoolExecutor(Plugin $plugin, string $autoload, string $hostname) : ThreadPoolExecutor {
		return new ThreadPoolExecutor(new ThreadFactory(
			Executor::class, LoggerUtils::makeLogger($plugin), $autoload,
			static function(Executor $e) use ($autoload, $hostname) : array {
				require_once $autoload;
				$client = new AnticheatClient($hostname, ['credentials' => ChannelCredentials::createInsecure()]);
				if (!$client->waitForReady(1000 * 10)) {
					throw new RuntimeException('failed to connect to server');
				}
				return [$client];
			},
			static fn(AnticheatClient $client) => $client->close()
		), 2);
	}

	/**
	 * @param JudgementData[] $judgementsList
	 */
	private function handleJudgements(Player $p, array $judgementsList) : void {
		foreach ($judgementsList as $j) {
			$formattedString = sprintf("judgement: %s: %s message:%s",
				$j->getType(),
				$j->getJudgement(),
				$j->getMessage()
			);
			switch ($j->getJudgement()) {
				case Judgement::DEBUG:
				case Judgement::AMBIGUOUS:
					$p->sendMessage($formattedString);
					break;
				case Judgement::TRIGGER:
					$p->kick($formattedString, false);
					break 2;
			}
		}
	}

	private function getAddPlayerRequest(Player $player) : AddPlayerRequest {
		return (new AddPlayerRequest())
			->setPlayer((new XyronPlayer())
				->setOs(Convert::deviceOS($player->getPlayerInfo()->getExtraData()["DeviceOS"]))
				->setName($player->getName())
			)->setData([(new TimestampedReportData())
				->setData([Convert::wildcard((new PlayerGameModeData())
						->setGameMode(Convert::gameMode($player->getGamemode()))
					), Convert::wildcard((new PlayerInputModeData())
						->setInputMode(Convert::inputMode(
							$player->getPlayerInfo()->getExtraData()["DeviceOS"]
						))
					), Convert::wildcard(
						$this->getEffectData($player)
					)]
				),
			]);
	}

	public function onPlayerInit(PlayerJoinEvent $ev) : void {
		$player = $ev->getPlayer();
		Await::do(function() use ($player) : void {
			$reqData = $this->getAddPlayerRequest($player)->serializeToString();

			$receipt = new PlayerReceipt();
			$receipt->mergeFromString(Await::fiberAsync(static function(AnticheatClient $client) use ($reqData) : string {
				$req = (new AddPlayerRequest());
				$req->mergeFromString($reqData);
				[$resp, $status] = $client->AddPlayer($req)->wait();
				if ($status->code !== STATUS_OK) {
					throw new RuntimeException($status);
				}
				assert($resp instanceof PlayerReceipt);
				return $resp->serializeToString();
			}, $this->executor));

			if (!$player->isOnline()) {
				return;
			}
			$data = new XyronData(
				$receipt,
				new BufferedDataQueue()
			);
			$this->data[spl_object_id($player)] = $data;
		})->panic();
	}

	public function onPlayerQuit(PlayerQuitEvent $ev) : void {
		$data = $this->data[spl_object_id($ev->getPlayer())] ?? null;
		if ($data !== null) {
			$receiptData = $data->receipt->serializeToString();
			Await::do(function() use ($receiptData) : void {
				Await::fiberAsync(static function(AnticheatClient $client) use ($receiptData) : void {
					$receipt = new PlayerReceipt();
					$receipt->mergeFromString($receiptData);
					[$resp, $status] = $client->RemovePlayer($receipt)->wait();
					if ($status->code !== STATUS_OK) {
						throw new RuntimeException($status);
					}
				}, $this->executor);
			})->panic();
		}
		unset($this->data[spl_object_id($ev->getPlayer())]);
	}

	public function onPlayerGameModeChange(PlayerGameModeChangeEvent $ev) : void {
		$player = $ev->getPlayer();
		$data = $this->data[spl_object_id($player)] ?? null;
		if ($data !== null) {
			$data->getQueue()->add($this->getTick(), Convert::wildcard((new PlayerGameModeData())
				->setGameMode(Convert::gameMode($player->getGamemode()))
			));
		}
	}


	/* Player Action Data Start */
	private function handleAction(Player $player, int $action) : void {
		$data = $this->data[spl_object_id($player)] ?? null;
		if ($data !== null) {
			$data->getQueue()->add($this->getTick(), Convert::wildcard(
				$this->getActionData($player, $action)
			));
		}
	}

	public function onPlayerToggleSprint(PlayerToggleSprintEvent $ev) : void {
		$player = $ev->getPlayer();
		$action = $ev->isSprinting() ?
			PlayerAction::StartSprint :
			PlayerAction::StopSprint;
		$this->handleAction($player, $action);
	}

	public function onPlayerToggleSneak(PlayerToggleSneakEvent $ev) : void {
		$player = $ev->getPlayer();
		$action = $ev->isSneaking() ?
			PlayerAction::StartSneak :
			PlayerAction::StopSneak;
		$this->handleAction($player, $action);
	}

	public function onPlayerToggleFlight(PlayerToggleFlightEvent $ev) : void {
		$player = $ev->getPlayer();
		$action = $ev->isFlying() ?
			PlayerAction::StartSprintFlying :
			PlayerAction::StopSprintFlying;
		$this->handleAction($player, $action);
	}

	public function onPlayerToggleGlide(PlayerToggleGlideEvent $ev) : void {
		$player = $ev->getPlayer();
		$action = $ev->isGliding() ?
			PlayerAction::StartGliding :
			PlayerAction::StopGliding;
		$this->handleAction($player, $action);
	}

	public function onPlayerToggleSwim(PlayerToggleSwimEvent $ev) : void {
		$player = $ev->getPlayer();
		$action = $ev->isSwimming() ?
			PlayerAction::StartSwimming :
			PlayerAction::StopSwimming;
		$this->handleAction($player, $action);
	}

	public function onInventoryOpen(InventoryOpenEvent $ev) : void {
		$player = $ev->getPlayer();
		$this->handleAction($player, PlayerAction::OpenInventory);
	}

	public function onInventoryClose(InventoryCloseEvent $ev) : void {
		$player = $ev->getPlayer();
		$this->handleAction($player, PlayerAction::CloseInventory);
	}

	/* Player Action Data End */

	/* Player Move Data Start */
	public function onPlayerMove(PlayerMoveEvent $ev) : void {
		$player = $ev->getPlayer();
		$data = $this->data[spl_object_id($player)] ?? null;
		if ($data !== null) {
			$data->getQueue()->add($this->getTick(), Convert::wildcard(
				$this->getMovementData($player, $ev->getTo(), false)
			));
		}
	}

	public function onPlayerTeleport(EntityTeleportEvent $ev) : void {
		if ($ev->getEntity() instanceof Player) {
			$player = $ev->getEntity();
			$data = $this->data[spl_object_id($player)] ?? null;
			if ($data !== null) {
				$data->getQueue()->add($this->getTick(), Convert::wildcard(
					$this->getMovementData($player, $ev->getTo(), true)
				));
			}
		}
	}

	/* Player Move Data End */
	/* PlaceBlock, BreakBlock Start */

	public function onBlockPlace(BlockPlaceEvent $ev) : void {
		$player = $ev->getPlayer();
		$data = $this->data[spl_object_id($player)] ?? null;
		if ($data !== null) {
			foreach ($ev->getTransaction()->getBlocks() as [$x, $y, $z, $block]) {
				$data->getQueue()->add($this->getTick(), Convert::wildcard((new PlayerPlaceBlockData())
					->setPlacedBlock(Convert::xyronBlock($block, new Vector3($x, $y, $z)))
					->setPosition($this->getPositionData($player, $player->getPosition()))
				));
			}
		}
	}

	public function onBlockBreak(BlockBreakEvent $ev) : void {
		$player = $ev->getPlayer();
		$data = $this->data[spl_object_id($player)] ?? null;
		if ($data !== null) {
			$data->getQueue()->add($this->getTick(), Convert::wildcard((new PlayerBreakBlockData())
				->setBrokenBlock(Convert::xyronBlockGetPosition($ev->getBlock()))
				->setPosition($this->getPositionData($player, $player->getPosition()))
			));
		}
	}

	/* PlaceBlock, BreakBlock End */

	public function onPlayerConsume(PlayerItemConsumeEvent $ev) : void {
		$player = $ev->getPlayer();
		$data = $this->data[spl_object_id($player)] ?? null;
		if ($data !== null) {
			$data->getQueue()->add($this->getTick(), Convert::wildcard(
				(new PlayerEatFoodData())
					//FIXME
					->setStatus(ConsumeStatus::Stop)
			));
		}
	}

	public function onPVP(EntityDamageByEntityEvent $ev) : void {
		$player = $ev->getEntity();
		$damager = $ev->getDamager();
		if ($player instanceof Player && $damager instanceof Player) {
			$data = $this->data[spl_object_id($player)] ?? null;
			if ($data !== null) {
				$data->getQueue()->add($this->getTick(), Convert::wildcard(
					(new PlayerAttackData())
						->setData((new AttackData())
							->setCause(Convert::damageCause($ev->getCause()))
							->setAttacker($this->getPositionData($damager, $damager->getPosition()))
							->setTarget($this->getPositionData($player, $player->getPosition()))
						)
				));
			}
		}
	}

	public function onPlayerJump(PlayerJumpEvent $ev) : void {
		$player = $ev->getPlayer();
		$data = $this->data[spl_object_id($player)] ?? null;
		if ($data !== null) {
			$data->getQueue()->add($this->getTick(), Convert::wildcard(
				$this->getActionData($player, PlayerAction::Jump)
			));
		}
	}

	public function onPlayerMotion(EntityMotionEvent $ev) : void {
		if ($ev->getEntity() instanceof Player) {
			$player = $ev->getEntity();
			$data = $this->data[spl_object_id($player)] ?? null;
			if ($data !== null) {
				$data->getQueue()->add($this->getTick(), Convert::wildcard(
					(new PlayerMotionData())
						->setMotion(Convert::xyronVec3f($ev->getVector()))
						->setPosition($this->getPositionData($player, $player->getPosition()))
				));
			}
		}
	}

	public function onPlayerDeath(PlayerDeathEvent $ev) : void {
		$player = $ev->getEntity();
		$data = $this->data[spl_object_id($player)] ?? null;
		if ($data !== null) {
			$data->getQueue()->add($this->getTick(), Convert::wildcard(
				$this->getLifeData(false)
			));
		}
	}

	public function onPlayerRespawn(PlayerRespawnEvent $ev) : void {
		$player = $ev->getPlayer();
		$data = $this->data[spl_object_id($player)] ?? null;
		if ($data !== null) {
			$data->getQueue()->add($this->getTick(), Convert::wildcard(
				$this->getLifeData(true)
			));
		}
	}

	public function onEffectAdd(EntityEffectAddEvent $ev) : void {
		$player = $ev->getEntity();
		if ($player instanceof Player) {
			$data = $this->data[spl_object_id($player)] ?? null;
			if ($data !== null) {
				$data->getQueue()->add($this->getTick(), Convert::wildcard(
					$this->getEffectData($player)
				));
			}
		}
	}

	public function onEffectRemove(EntityEffectRemoveEvent $ev) : void {
		$player = $ev->getEntity();
		if ($player instanceof Player) {
			$data = $this->data[spl_object_id($player)] ?? null;
			if ($data !== null) {
				$data->getQueue()->add($this->getTick(), Convert::wildcard(
					$this->getEffectData($player)
				));
			}
		}
	}

	private function getActionData(Player $player, int $action) : PlayerActionData {
		return (new PlayerActionData())
			->setAction($action)
			->setPosition($this->getPositionData($player, $player->getPosition()));
	}

	private function getLifeData(bool $alive) : PlayerLifeData {
		return (new PlayerLifeData())
			->setAlive($alive);
	}

	private function getMovementData(Player $player, Vector3 $to, bool $teleport) : PlayerMoveData {
		return (new PlayerMoveData())
			->setNewPosition($this->getPositionData($player, $to))
			->setTeleport($teleport);
	}

	private function getIntersectedBlock(World $world, AxisAlignedBB $bb) : \Generator {
		$inset = 0.002;
		$minX = (int) floor($bb->minX - $inset);
		$minY = (int) floor($bb->minY - $inset);
		$minZ = (int) floor($bb->minZ - $inset);
		$maxX = (int) floor($bb->maxX + $inset);
		$maxY = (int) floor($bb->maxY + $inset);
		$maxZ = (int) floor($bb->maxZ + $inset);

		for ($z = $minZ; $z <= $maxZ; ++$z) {
			for ($x = $minX; $x <= $maxX; ++$x) {
				for ($y = $minY; $y <= $maxY; ++$y) {
					$blk = $world->getBlockAt($x, $y, $z);
					if(!$blk instanceof Air){
						yield $blk;
					}
				}
			}
		}
	}

	private function wouldCollideVertically(Player $player, Vector3 $newPos) : bool {
		$bb = clone $player->getBoundingBox();

		$xLen = $bb->getXLength();
		$yLen = $bb->getYLength();
		$zLen = $bb->getZLength();

		$oldPos = $player->getPosition();
		$dx = $newPos->getX() - $oldPos->getX();
		$dy = $newPos->getY() - $oldPos->getY();
		$dz = $newPos->getZ() - $oldPos->getZ();

		if (abs($dx) <= $xLen && abs($dy) <= $yLen && abs($dz) <= $zLen) {
			$bb = $bb->addCoord($dx, $dy, $dz);
			return count($player->getWorld()->getCollisionBlocks($bb, true)) !== 0;
		}

		return false;
	}

	private function getPositionData(Player $player, Vector3 $newPos) : EntityPositionData {
		$newPosBB = clone $player->getBoundingBox();
		$delta = $newPos->subtractVector($player->getPosition());
		$newPosBB = $newPosBB->addCoord($delta->getX(), $delta->getY(), $delta->getZ());

		$blockPos = new Vector3($newPos->x, $newPosBB->minY, $newPos->z);
		$below = $player->getWorld()->getBlock($blockPos);
		$collision = array_map(
			Convert::xyronBlockGetPosition(...),
			$player->getWorld()->getCollisionBlocks($newPosBB)
		);

		$intersected = array_map(
			Convert::xyronBlockGetPosition(...),
			iterator_to_array($this->getIntersectedBlock($player->getWorld(), $newPosBB))
		);

		return (new EntityPositionData())
			->setPosition(Convert::xyronVec3f($player->getPosition()))
			->setDirection(Convert::xyronVec3f($player->getDirectionVector()))
			->setBoundingBox(Convert::xyronBoundingBox($player->getBoundingBox()))
			->setIsImmobile($player->hasNoClientPredictions())
			->setIsOnGround($player->isOnGround())
			->setAllowFlying($player->getAllowFlight())
			->setIsFlying($player->isFlying())
			//TODO improve this
			->setHaveGravity(true)
			->setMovementSpeed($player->getMovementSpeed())
			->setWouldCollideVertically($this->wouldCollideVertically($player, $newPos))
			->setBelowThatAffectMovement(Convert::xyronBlock($below, $blockPos))
			->setCollidedBlocks($collision)
			->setIntersectedBlocks($intersected);
	}

	private function getTick() : int {
		return Server::getInstance()->getTick();
	}

	private function getEffectData(Player $player) : PlayerEffectData {
		return (new PlayerEffectData())
			->setEffect(Convert::effects($player->getEffects()));
	}
}