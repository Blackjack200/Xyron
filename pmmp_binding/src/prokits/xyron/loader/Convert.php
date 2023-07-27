<?php

namespace prokits\xyron\loader;

use Closure;
use pocketmine\block\Block;
use pocketmine\block\BlockTypeIds;
use pocketmine\block\Liquid;
use pocketmine\entity\effect\Effect;
use pocketmine\entity\effect\EffectManager;
use pocketmine\entity\effect\VanillaEffects;
use pocketmine\event\entity\EntityDamageEvent;
use pocketmine\item\Armor;
use pocketmine\item\Bow;
use pocketmine\item\Durable;
use pocketmine\item\Food;
use pocketmine\item\Item;
use pocketmine\item\ItemTypeIds;
use pocketmine\item\ProjectileItem;
use pocketmine\item\Tool;
use pocketmine\math\AxisAlignedBB;
use pocketmine\math\Vector3;
use pocketmine\player\GameMode;
use prokits\xyron\AxisAlignedBoundingBox;
use prokits\xyron\BlockData;
use prokits\xyron\BlockFeature;
use prokits\xyron\DamageCause;
use prokits\xyron\EffectFeature;
use prokits\xyron\GameMode as XyronGameMode;
use prokits\xyron\ItemData;
use prokits\xyron\ItemFeature;
use prokits\xyron\PlayerActionData;
use prokits\xyron\PlayerAttackData;
use prokits\xyron\PlayerBreakBlockData;
use prokits\xyron\PlayerEatFoodData;
use prokits\xyron\PlayerEffectData;
use prokits\xyron\PlayerGameModeData;
use prokits\xyron\PlayerHeldItemChangeData;
use prokits\xyron\PlayerInputModeData;
use prokits\xyron\PlayerLifeData;
use prokits\xyron\PlayerMotionData;
use prokits\xyron\PlayerMoveData;
use prokits\xyron\PlayerPlaceBlockData;
use prokits\xyron\Vec3f;
use prokits\xyron\Vec3i;
use prokits\xyron\WildcardReportData;

final class Convert {
	private function __construct() { }

	public static function deviceOS(int $os) : int {
		return $os;
	}

	public static function gameMode(GameMode $gameMode) : int {
		return match ($gameMode) {
			GameMode::CREATIVE() => XyronGameMode::Creative,
			GameMode::ADVENTURE() => XyronGameMode::Adventure,
			GameMode::SURVIVAL() => XyronGameMode::Survival,
			GameMode::SPECTATOR() => XyronGameMode::Spectator,
		};
	}

	public static function damageCause(int $cause) : int {
		return match ($cause) {
			EntityDamageEvent::CAUSE_CONTACT => DamageCause::Contact,
			EntityDamageEvent::CAUSE_ENTITY_ATTACK => DamageCause::EntityAttack,
			EntityDamageEvent::CAUSE_PROJECTILE => DamageCause::Projectile,
			EntityDamageEvent::CAUSE_SUFFOCATION => DamageCause::Suffocation,
			EntityDamageEvent::CAUSE_FALL => DamageCause::Fall,
			EntityDamageEvent::CAUSE_FIRE => DamageCause::Fire,
			EntityDamageEvent::CAUSE_FIRE_TICK => DamageCause::FireTick,
			EntityDamageEvent::CAUSE_LAVA => DamageCause::Lava,
			EntityDamageEvent::CAUSE_DROWNING => DamageCause::Drowning,
			EntityDamageEvent::CAUSE_BLOCK_EXPLOSION => DamageCause::BlockExplosion,
			EntityDamageEvent::CAUSE_ENTITY_EXPLOSION => DamageCause::EntityExplosion,
			EntityDamageEvent::CAUSE_VOID => DamageCause::Void,
			EntityDamageEvent::CAUSE_SUICIDE => DamageCause::Suicide,
			EntityDamageEvent::CAUSE_MAGIC => DamageCause::Magic,
			EntityDamageEvent::CAUSE_CUSTOM => DamageCause::Custom,
			EntityDamageEvent::CAUSE_STARVATION => DamageCause::Starvation
		};
	}

	public static function inputMode(int $mode) : int {
		return $mode;
	}

	public static function xyronVec3f(Vector3 $vec3) : Vec3f {
		return (new Vec3f())
			->setX((float) $vec3->getX())
			->setY((float) $vec3->getY())
			->setZ((float) $vec3->getZ());
	}

	public static function vec3f(Vec3f $v3f) : Vector3 {
		return new Vector3($v3f->getX(), $v3f->getY(), $v3f->getZ());
	}


	public static function xyronVec3i(Vector3 $vec3) : Vec3i {
		return (new Vec3i())
			->setX((int) $vec3->getX())
			->setY((int) $vec3->getY())
			->setZ((int) $vec3->getZ());
	}

	public static function vec3i(Vec3f $v3i) : Vector3 {
		return new Vector3($v3i->getX(), $v3i->getY(), $v3i->getZ());
	}

	public static function xyronBoundingBox(AxisAlignedBB $aabb) : AxisAlignedBoundingBox {
		return (new AxisAlignedBoundingBox())
			->setMin(self::xyronVec3f(new Vector3($aabb->minX, $aabb->minY, $aabb->minZ)))
			->setMax(self::xyronVec3f(new Vector3($aabb->maxX, $aabb->maxY, $aabb->maxZ)));
	}

	public static function boundingBox(AxisAlignedBoundingBox $aabb) : AxisAlignedBB {
		$min = $aabb->getMin();
		$max = $aabb->getMax();
		if ($min === null || $max === null) {
			throw new \RuntimeException();
		}
		return new AxisAlignedBB(
			$min->getX(), $min->getY(), $min->getZ(),
			$max->getX(), $max->getY(), $max->getZ(),
		);
	}

	public static function xyronBlockGetPosition(Block $blk) : BlockData {
		return self::xyronBlock($blk, $blk->getPosition());
	}

	public static function xyronBlock(Block $blk, Vector3 $position) : BlockData {
		$boxes = [];
		foreach ($blk->getCollisionBoxes() as $box) {
			$boxes[] = self::xyronBoundingBox($box);
		}
		return (new BlockData())
			->setPosition(self::xyronVec3i($position))
			->setFeature((new BlockFeature())
				->setCollisionBoxes($boxes)
				->setFriction((float) $blk->getFrictionFactor())
				->setIsSolid($blk->isSolid())
				->setIsLiquid($blk instanceof Liquid)
				->setIsAir($blk->getTypeId() === BlockTypeIds::AIR)
				->setIsSlime($blk->getTypeId() === BlockTypeIds::SLIME)
				->setIsClimbable($blk->canClimb())
				->setIsIce($blk->getTypeId() === BlockTypeIds::ICE)
				->setIsCobweb($blk->getTypeId() === BlockTypeIds::COBWEB)
				->setIsSweetBerry($blk->getTypeId() === BlockTypeIds::SWEET_BERRY_BUSH)
			);
	}

	public static function xyronItem(Item $item) : ItemData {
		return (new ItemData())
			->setVanillaName($item->getName())
			->setCount($item->getCount())
			->setFeature((new ItemFeature())
				->setIsArmor($item instanceof Armor)
				//TODO figure out what is this
				->setIsBlockPlanterItem(false)
				->setIsDamageable($item instanceof Durable)
				->setIsFood($item instanceof Food)
				->setIsThrowable($item instanceof ProjectileItem)
				->setIsTool($item instanceof Tool)
				->setIsBow($item instanceof Bow || $item->getTypeId() === ItemTypeIds::BOW)
				->setIsCrossBow($item->getTypeId() === ItemTypeIds::CROSSBOW)
				->setIsShield($item->getTypeId() === ItemTypeIds::SHIELD)
			);
	}

	/**
	 * @return EffectFeature[]
	 */
	public static function effects(EffectManager $effects) : array {
		$l = [];
		self::internalEffect($l, $effects, VanillaEffects::SPEED(), static fn(EffectFeature $e) => $e->setIsSpeed(true));
		self::internalEffect($l, $effects, VanillaEffects::HASTE(), static fn(EffectFeature $e) => $e->setIsHaste(true));
		//self::internalEffect($l, $effects, VanillaEffects::SLOW_FALLING(), static fn(EffectFeature $e) => $e->setIsSlowFalling(true));
		self::internalEffect($l, $effects, VanillaEffects::LEVITATION(), static fn(EffectFeature $e) => $e->setIsLevitation(true));
		self::internalEffect($l, $effects, VanillaEffects::SLOWNESS(), static fn(EffectFeature $e) => $e->setIsSlowness(true));
		self::internalEffect($l, $effects, VanillaEffects::JUMP_BOOST(), static fn(EffectFeature $e) => $e->setIsJumpBoost(true));
		return $l;
	}

	private static function internalEffect(array &$ls, EffectManager $effects, Effect $effect, Closure $c) : void {
		$eff = $effects->get($effect);
		if ($eff !== null) {
			$ef = (new EffectFeature())
				->setAmplifier($eff->getAmplifier());
			$c($ef);
			$ls[] = $ef;
		}
	}

	/**
	 * @param PlayerActionData|PlayerMoveData|PlayerPlaceBlockData|PlayerBreakBlockData|PlayerEatFoodData|PlayerAttackData|PlayerEffectData|PlayerGameModeData|PlayerMotionData|PlayerInputModeData|PlayerHeldItemChangeData|PlayerLifeData $data
	 */
	public static function wildcard($data) : WildcardReportData {
		if ($data instanceof PlayerActionData) {
			return (new WildcardReportData())->setActionData($data);
		}

		if ($data instanceof PlayerMoveData) {
			return (new WildcardReportData())->setMoveData($data);
		}

		if ($data instanceof PlayerPlaceBlockData) {
			return (new WildcardReportData())->setPlaceBlockData($data);
		}

		if ($data instanceof PlayerBreakBlockData) {
			return (new WildcardReportData())->setBreakBlockData($data);
		}

		if ($data instanceof PlayerEatFoodData) {
			return (new WildcardReportData())->setEatFoodData($data);
		}

		if ($data instanceof PlayerAttackData) {
			return (new WildcardReportData())->setAttackData($data);
		}

		if ($data instanceof PlayerEffectData) {
			return (new WildcardReportData())->setEffectData($data);
		}

		if ($data instanceof PlayerGameModeData) {
			return (new WildcardReportData())->setGameModeData($data);
		}

		if ($data instanceof PlayerMotionData) {
			return (new WildcardReportData())->setMotionData($data);
		}

		if ($data instanceof PlayerInputModeData) {
			return (new WildcardReportData())->setInputModeData($data);
		}

		if ($data instanceof PlayerHeldItemChangeData) {
			return (new WildcardReportData())->setHeldItemChangeData($data);
		}

		if ($data instanceof PlayerLifeData) {
			return (new WildcardReportData())->setLifeData($data);
		}
		throw new \RuntimeException();
	}
}