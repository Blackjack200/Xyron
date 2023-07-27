package com.blackjack200.xyron.nukkit;

import cn.nukkit.AdventureSettings;
import cn.nukkit.Player;
import cn.nukkit.Server;
import cn.nukkit.block.Block;
import cn.nukkit.event.EventHandler;
import cn.nukkit.event.Listener;
import cn.nukkit.event.block.BlockBreakEvent;
import cn.nukkit.event.block.BlockPlaceEvent;
import cn.nukkit.event.entity.EntityDamageByEntityEvent;
import cn.nukkit.event.entity.EntityMotionEvent;
import cn.nukkit.event.entity.EntityTeleportEvent;
import cn.nukkit.event.player.*;
import cn.nukkit.level.Level;
import cn.nukkit.level.Position;
import cn.nukkit.math.AxisAlignedBB;
import cn.nukkit.math.Vector3;
import cn.nukkit.plugin.PluginBase;
import com.github.blackjack200.xyron.*;
import com.google.common.collect.Lists;
import io.grpc.ManagedChannelBuilder;
import lombok.val;
import lombok.var;

import java.util.Arrays;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.stream.Collectors;

public class Loader extends PluginBase implements Listener {
    private AnticheatGrpc.AnticheatFutureStub client;
    private final BufferedDataFlushPool pool = new BufferedDataFlushPool();
    private final Map<Player, XyronData> data = new HashMap<>(32);

    @Override
    @lombok.SneakyThrows
    public void onEnable() {
        this.getLogger().info("Connecting to the anticheat...");
        val channel = ManagedChannelBuilder.forAddress("localhost", 8884).usePlaintext().build();
        this.client = AnticheatGrpc.newFutureStub(channel).withWaitForReady();
        this.getLogger().info("Connected to the anticheat");

        this.getServer().getPluginManager().registerEvents(this, this);
        this.getServer().getScheduler().scheduleRepeatingTask(this, () -> {
            this.data.forEach((p, x) -> {
                x.getQueue().add(getTick(), Convert.wildcard(
                        getEffectData(p)
                ));
                this.pool.add(
                        x.getQueue().flush(this.client, x.getReceipt(), this.getServer().getTick(), p.getPing() / 1000D),
                        (resp) -> {
                            if (!p.isOnline()) {
                                //the player has been quit
                                return;
                            }
                            this.handleJudgements(p, resp.getJudgementsList());
                        }
                );
            });
            this.pool.poll();
        }, 10);
        val req = Xchange.AddPlayerRequest.newBuilder()
                .setPlayer(PlayerOuterClass.Player.newBuilder()
                        .setOsValue(PlayerOuterClass.DeviceOS.Android_VALUE)
                        .setName("IPlayfordev")
                );
        req.putData(0L, Xchange.TimestampedReportData.newBuilder()
                .addData(PlayerWrappers.WildcardReportData.newBuilder().setGameModeData(
                        PlayerWrappers.PlayerGameModeData.newBuilder()
                                .setGameModeValue(PlayerOuterClass.GameMode.Survival_VALUE)
                )).build()
        );
    }

    @Override
    public void onDisable() {
        this.getLogger().info("Anticheat closing...");
        this.pool.shutdown();
        this.getLogger().info("Anticheat closed");
    }

    private void handleJudgements(Player p, List<Xchange.JudgementData> judgementsList) {
        for (val j : judgementsList) {
            val formattedString = String.format("judgement: %s: %s message:%s",
                    j.getType(),
                    j.getJudgement(),
                    j.getMessage()
            );
            switch (j.getJudgement().getNumber()) {
                case AnticheatTypes.Judgement.DEBUG_VALUE:
                case AnticheatTypes.Judgement.AMBIGUOUS_VALUE:
                    p.sendMessage(formattedString);
                    break;
                case AnticheatTypes.Judgement.TRIGGER_VALUE:
                    p.kick(formattedString, false);
                    break;
            }
        }
    }

    @EventHandler
    public void onPlayerInit(PlayerLocallyInitializedEvent ev) {
        val player = ev.getPlayer();
        player.setCheckMovement(false);

        val req = Xchange.AddPlayerRequest.newBuilder()
                .setPlayer(PlayerOuterClass.Player.newBuilder()
                        .setOs(Convert.deviceOS(player.getLoginChainData().getDeviceOS()))
                        .setName(player.getName())
                );
        req.putData(0L, Xchange.TimestampedReportData.newBuilder()
                .addData(Convert.wildcard(
                        PlayerWrappers.PlayerGameModeData.newBuilder()
                                .setGameMode(Convert.gameMode(player.getGamemode()))
                                .build()
                ))
                .addData(Convert.wildcard(
                        PlayerWrappers.PlayerInputModeData.newBuilder()
                                .setInputMode(Convert.inputMode(
                                        player.getLoginChainData().getCurrentInputMode()
                                ))
                                .build()
                ))
                .addData(Convert.wildcard(
                        getEffectData(player)
                ))
                .build()
        );

        this.pool.add(this.client.addPlayer(req.build()), (receipt) -> {
            if (!player.isOnline()) {
                return;
            }
            val data = new XyronData(
                    receipt,
                    new BufferedDataQueue()
            );
            this.data.put(player, data);
        });
    }

    @EventHandler
    public void onPlayerQuit(PlayerQuitEvent ev) {
        val data = this.data.get(ev.getPlayer());
        if (data != null) {
            this.pool.add(this.client.removePlayer(data.getReceipt()), (e) -> {
            });
        }
        this.data.remove(ev.getPlayer());
    }

    @EventHandler
    public void onPlayerGameModeChange(PlayerGameModeChangeEvent ev) {
        val player = ev.getPlayer();
        val data = this.data.get(player);
        if (data != null) {
            data.getQueue().add(getTick(), Convert.wildcard(
                    PlayerWrappers.PlayerGameModeData.newBuilder()
                            .setGameMode(Convert.gameMode(player.getGamemode()))
                            .build()
            ));
        }
    }


    /* Player Action Data Start */
    private void handleAction(Player player, PlayerOuterClass.PlayerAction action) {
        val data = this.data.get(player);
        if (data != null) {
            data.getQueue().add(getTick(), Convert.wildcard(
                    getActionData(player, action)
            ));
        }
    }

    @EventHandler
    public void onPlayerToggleSprint(PlayerToggleSprintEvent ev) {
        val player = ev.getPlayer();
        val action = ev.isSprinting() ?
                PlayerOuterClass.PlayerAction.StartSprint :
                PlayerOuterClass.PlayerAction.StopSprint;
        this.handleAction(player, action);
    }

    @EventHandler
    public void onPlayerToggleSneak(PlayerToggleSneakEvent ev) {
        val player = ev.getPlayer();
        val action = ev.isSneaking() ?
                PlayerOuterClass.PlayerAction.StartSneak :
                PlayerOuterClass.PlayerAction.StopSneak;
        this.handleAction(player, action);
    }

    @EventHandler
    public void onPlayerToggleFlight(PlayerToggleFlightEvent ev) {
        val player = ev.getPlayer();
        val action = ev.isFlying() ?
                PlayerOuterClass.PlayerAction.StartSprintFlying :
                PlayerOuterClass.PlayerAction.StopSprintFlying;
        this.handleAction(player, action);
    }

    @EventHandler
    public void onPlayerToggleGlide(PlayerToggleGlideEvent ev) {
        val player = ev.getPlayer();
        val action = ev.isGliding() ?
                PlayerOuterClass.PlayerAction.StartGliding :
                PlayerOuterClass.PlayerAction.StopGliding;
        this.handleAction(player, action);
    }

    @EventHandler
    public void onPlayerToggleSwim(PlayerToggleSwimEvent ev) {
        val player = ev.getPlayer();
        val action = ev.isSwimming() ?
                PlayerOuterClass.PlayerAction.StartSwimming :
                PlayerOuterClass.PlayerAction.StopSwimming;
        this.handleAction(player, action);
    }
    /* Player Action Data End */

    /* Player Move Data Start */
    @EventHandler
    public void onPlayerMove(PlayerMoveEvent ev) {
        val player = ev.getPlayer();
        val data = this.data.get(player);
        if (data != null) {
            data.getQueue().add(getTick(), Convert.wildcard(
                    getMovementData(player, ev.getTo(), false)
            ));
        }
    }

    @EventHandler
    public void onPlayerTeleport(EntityTeleportEvent ev) {
        if (ev.getEntity() instanceof Player) {
            val player = (Player) ev.getEntity();
            val data = this.data.get(player);
            if (data != null) {
                data.getQueue().add(getTick(), Convert.wildcard(
                        getMovementData(player, ev.getTo(), true)
                ));
            }
        }
    }

    /* Player Move Data End */
    /* PlaceBlock, BreakBlock Start */

    @EventHandler
    public void onBlockPlace(BlockPlaceEvent ev) {
        val player = ev.getPlayer();
        val data = this.data.get(player);
        if (data != null) {
            data.getQueue().add(getTick(), Convert.wildcard(
                    PlayerWrappers.PlayerPlaceBlockData.newBuilder()
                            .setPlacedBlock(Convert.block(ev.getBlock()))
                            .setPosition(getPositionData(player))
                            .build()
            ));
        }
    }

    @EventHandler
    public void onBlockBreak(BlockBreakEvent ev) {
        val player = ev.getPlayer();
        val data = this.data.get(player);
        if (data != null) {
            data.getQueue().add(getTick(), Convert.wildcard(
                    PlayerWrappers.PlayerBreakBlockData.newBuilder()
                            .setBrokenBlock(Convert.block(ev.getBlock()))
                            .setPosition(getPositionData(player))
                            .build()
            ));
        }
    }
    /* PlaceBlock, BreakBlock End */

    @EventHandler
    public void onPlayerConsume(PlayerItemConsumeEvent ev) {
        val player = ev.getPlayer();
        val data = this.data.get(player);
        if (data != null) {
            data.getQueue().add(getTick(), Convert.wildcard(
                    PlayerWrappers.PlayerEatFoodData.newBuilder()
                            //FIXME
                            .setStatus(PlayerOuterClass.ConsumeStatus.Stop)
                            .build()
            ));
        }
    }

    @EventHandler
    public void onPVP(EntityDamageByEntityEvent ev) {
        if (ev.getEntity() instanceof Player && ev.getDamager() instanceof Player) {
            val player = (Player) ev.getEntity();
            val damager = (Player) ev.getDamager();
            val data = this.data.get(player);
            if (data != null) {
                data.getQueue().add(getTick(), Convert.wildcard(
                        PlayerWrappers.PlayerAttackData.newBuilder()
                                .setData(PlayerOuterClass.AttackData.newBuilder()
                                        .setCause(Convert.damageCause(ev.getCause()))
                                        .setAttacker(getPositionData(damager))
                                        .setTarget(getPositionData(player))
                                )
                                .build()
                ));
            }
        }
    }

    @EventHandler
    public void onPlayerJump(PlayerJumpEvent ev) {
        val player = ev.getPlayer();
        val data = this.data.get(player);
        if (data != null) {
            data.getQueue().add(getTick(), Convert.wildcard(
                    getActionData(player, PlayerOuterClass.PlayerAction.Jump)
            ));
        }
    }

    @EventHandler
    public void onPlayerMotion(EntityMotionEvent ev) {
        if (ev.getEntity() instanceof Player) {
            val player = (Player) ev.getEntity();
            val data = this.data.get(player);
            if (data != null) {
                data.getQueue().add(getTick(), Convert.wildcard(
                        PlayerWrappers.PlayerMotionData.newBuilder()
                                .setMotion(Convert.vec3f(ev.getMotion()))
                                .setPosition(getPositionData(player))
                                .build()
                ));
            }
        }
    }

    @EventHandler
    public void onPlayerDeath(PlayerDeathEvent ev) {
        val player = ev.getEntity();
        val data = this.data.get(player);
        if (data != null) {
            data.getQueue().add(getTick(), Convert.wildcard(
                    getLifeData(false)
            ));
        }
    }

    @EventHandler
    public void onPlayerRespawn(PlayerRespawnEvent ev) {
        val player = ev.getPlayer();
        val data = this.data.get(player);
        if (data != null) {
            data.getQueue().add(getTick(), Convert.wildcard(
                    getLifeData(true)
            ));
        }
    }

    private PlayerWrappers.PlayerActionData getActionData(Player player, PlayerOuterClass.PlayerAction action) {
        return PlayerWrappers.PlayerActionData.newBuilder()
                .setAction(action)
                .setPosition(getPositionData(player))
                .build();
    }

    private PlayerWrappers.PlayerLifeData getLifeData(boolean alive) {
        return PlayerWrappers.PlayerLifeData.newBuilder()
                .setAlive(alive)
                .build();
    }

    private PlayerWrappers.PlayerMoveData getMovementData(Player player, Vector3 to, boolean teleport) {
        return PlayerWrappers.PlayerMoveData.newBuilder()
                .setNewPosition(getPositionData(player, to))
                .setTeleport(teleport)
                .build();
    }

    private PlayerOuterClass.EntityPositionData getPositionData(Player player) {
        return getPositionData(player, player.getPosition());
    }


    private List<Block> getIntersectedBlock(Level level, AxisAlignedBB bb) {
        List<Block> list = Lists.newLinkedList();

        int minX = (int) Math.floor(bb.getMinX() - 1);
        int minY = (int) Math.floor(bb.getMinY() - 1);
        int minZ = (int) Math.floor(bb.getMinZ() - 1);
        int maxX = (int) Math.floor(bb.getMaxX() + 1);
        int maxY = (int) Math.floor(bb.getMaxY() + 1);
        int maxZ = (int) Math.floor(bb.getMaxZ() + 1);

        for (int x = minX; x <= maxX; ++x) {
            for (int y = minY; y <= maxY; ++y) {
                for (int z = minZ; z <= maxZ; ++z) {
                    var block = level.getBlock(x, y, z);
                    var b = block.getCollisionBoundingBox();
                    if (b != null && b.intersectsWith(bb)) {
                        list.add(block);
                    }
                }
            }
        }
        return list;
    }

    private boolean wouldCollideVertically(Player player, Vector3 newPos) {
        AxisAlignedBB bb = player.getBoundingBox().clone();

        double xLen = bb.getMaxX() - bb.getMinX();
        double yLen = bb.getMaxY() - bb.getMinY();
        double zLen = bb.getMaxZ() - bb.getMinZ();

        Position oldPos = player.getPosition();
        double dx = newPos.getX() - oldPos.getX();
        double dy = newPos.getY() - oldPos.getY();
        double dz = newPos.getZ() - oldPos.getZ();

        if (Math.abs(dx) <= xLen && Math.abs(dy) <= yLen && Math.abs(dz) <= zLen) {
            bb = bb.addCoord(dx, dy, dz);
            return player.getLevel().getCollisionBlocks(bb, true).length == 1;
        } else {
            return false;
        }
    }

    private PlayerOuterClass.EntityPositionData getPositionData(Player player, Vector3 newPos) {
        var newPosBB = player.getBoundingBox().clone();
        val delta = newPos.subtract(player.getPosition());
        newPosBB = newPosBB.getOffsetBoundingBox(delta.getX(), delta.getY(), delta.getZ());
        newPosBB.setMinY(newPosBB.getMinY() - 0.50001);

        var below = player.getLevel().getBlock(new Vector3(newPos.x, newPosBB.getMinY(), newPos.z));

        var collision = Arrays.stream(player.getLevel().getCollisionBlocks(newPosBB))
                .map(Convert::block)
                .collect(Collectors.toList());

        var intersected = this.getIntersectedBlock(player.getLevel(), newPosBB)
                .stream()
                .map(Convert::block)
                .collect(Collectors.toList());

        return PlayerOuterClass.EntityPositionData.newBuilder()
                .setPosition(Convert.vec3f(player.getPosition()))
                .setDirection(Convert.vec3f(player.getDirectionVector()))
                .setBoundingBox(Convert.boundingBox(player.getBoundingBox()))
                .setIsImmobile(player.isImmobile())
                .setIsOnGround(player.isOnGround())
                .setAllowFlying(player.getAdventureSettings().get(AdventureSettings.Type.ALLOW_FLIGHT))
                .setIsFlying(player.getAdventureSettings().get(AdventureSettings.Type.FLYING))
                //TODO improve this
                .setHaveGravity(true)
                .setMovementSpeed(player.getMovementSpeed())
                .setWouldCollideVertically(this.wouldCollideVertically(player, newPos))
                .setBelowThatAffectMovement(Convert.block(below))
                .addAllCollidedBlocks(collision)
                .addAllIntersectedBlocks(intersected)
                .build();
    }

    private int getTick() {
        return Server.getInstance().getTick();
    }

    private PlayerWrappers.PlayerEffectData getEffectData(Player player) {
        return PlayerWrappers.PlayerEffectData.newBuilder().addAllEffect(Convert.effects(player.getEffects())).build();
    }
}
