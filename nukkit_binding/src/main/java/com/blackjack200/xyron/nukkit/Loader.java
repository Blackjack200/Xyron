package com.blackjack200.xyron.nukkit;

import cn.nukkit.AdventureSettings;
import cn.nukkit.Player;
import cn.nukkit.Server;
import cn.nukkit.block.Block;
import cn.nukkit.block.BlockID;
import cn.nukkit.block.BlockLiquid;
import cn.nukkit.event.EventHandler;
import cn.nukkit.event.Listener;
import cn.nukkit.event.entity.EntityTeleportEvent;
import cn.nukkit.event.player.*;
import cn.nukkit.math.AxisAlignedBB;
import cn.nukkit.math.Vector3;
import cn.nukkit.plugin.PluginBase;
import cn.nukkit.potion.Effect;
import com.github.blackjack200.xyron.*;
import com.google.common.collect.Lists;
import com.google.protobuf.Empty;
import io.grpc.ManagedChannelBuilder;
import lombok.val;
import lombok.var;

import java.util.*;
import java.util.function.Consumer;
import java.util.stream.Collectors;

public class Loader extends PluginBase implements Listener {
    private AnticheatGrpc.AnticheatFutureStub client;
    private final BufferedDataFlushPool<Xchange.ReportResponse> reportPool = new BufferedDataFlushPool<>();
    private final BufferedDataFlushPool<Xchange.PlayerReceipt> addPool = new BufferedDataFlushPool<>();
    private final BufferedDataFlushPool<Empty> removePool = new BufferedDataFlushPool<>();
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
            this.data.forEach((p, x) -> this.reportPool.add(
                    x.getQueue().flush(this.client, x.getReceipt(), this.getServer().getTick(), p.getPing() / 1000D),
                    (resp) -> {
                        if (!p.isOnline()) {
                            //the player has been quit
                            return;
                        }
                        for (val j : resp.getJudgementsList()) {
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
            ));
            this.addPool.poll();
            this.reportPool.poll();
            this.removePool.poll();
        }, 10);
    }

    @Override
    public void onDisable() {
        this.getLogger().info("Anticheat closing...");
        this.addPool.shutdown();
        this.reportPool.shutdown();
        this.removePool.shutdown();
        this.getLogger().info("Anticheat closed");
    }

    private PlayerOuterClass.DeviceOS convertDeviceOS(int os) {
        var c = PlayerOuterClass.DeviceOS.forNumber(os);
        if (c == null) {
            c = PlayerOuterClass.DeviceOS.Android;
        }
        return c;
    }

    @EventHandler
    public void onPlayerInit(PlayerLocallyInitializedEvent ev) {
        val player = ev.getPlayer();
        player.setCheckMovement(false);
        val req = Xchange.AddPlayerRequest.newBuilder()
                .setPlayer(PlayerOuterClass.Player.newBuilder()
                        .setOs(convertDeviceOS(player.getLoginChainData().getDeviceOS()))
                        .setName(player.getName())
                );
        req.putData(0L, Xchange.TimestampedReportData.newBuilder()
                .addData(PlayerWrappers.WildcardReportData.newBuilder().setGameModeData(
                        PlayerWrappers.PlayerGameModeData.newBuilder()
                                .setGameModeValue(PlayerOuterClass.GameMode.Survival_VALUE)
                ))
                .build()
        );
        this.addPool.add(this.client.addPlayer(req.build()), (receipt) -> {
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
            this.removePool.add(this.client.removePlayer(data.getReceipt()), (e) -> {
            });
        }
        this.data.remove(ev.getPlayer());
    }

    private PrimitiveTypes.Vec3f toVec3f(Vector3 vec3) {
        return PrimitiveTypes.Vec3f.newBuilder()
                .setX((float) vec3.getX())
                .setY((float) vec3.getY())
                .setZ((float) vec3.getZ())
                .build();
    }

    private PrimitiveTypes.Vec3i toVec3i(Vector3 vec3) {
        return PrimitiveTypes.Vec3i.newBuilder()
                .setX((int) vec3.getX())
                .setY((int) vec3.getY())
                .setZ((int) vec3.getZ())
                .build();
    }

    private Vector3 toVec3f(PrimitiveTypes.Vec3f v3f) {
        return new Vector3(v3f.getX(), v3f.getY(), v3f.getZ());
    }

    private PrimitiveTypes.AxisAlignedBoundingBox toBB(AxisAlignedBB aabb) {
        return PrimitiveTypes.AxisAlignedBoundingBox.newBuilder()
                .setMin(toVec3f(new Vector3(aabb.getMinX(), aabb.getMinY(), aabb.getMinZ())))
                .setMax(toVec3f(new Vector3(aabb.getMaxX(), aabb.getMaxY(), aabb.getMaxZ())))
                .build();
    }

    private PrimitiveTypes.BlockData toBlock(Block blk) {
        return PrimitiveTypes.BlockData.newBuilder()
                .setPosition(toVec3i(blk))
                .setFeature(PrimitiveTypes.BlockFeature.newBuilder()
                        .addAllCollisionBoxes(Lists.newArrayList(toBB(blk.getCollisionBoundingBox())))
                        .setFriction((float) blk.getFrictionFactor())
                        .setIsSolid(blk.isSolid())
                        .setIsLiquid(blk instanceof BlockLiquid)
                        .setIsAir(blk.getId() == BlockID.AIR)
                        .setIsSlime(blk.getId() == BlockID.SLIME_BLOCK)
                        .setIsClimbable(blk.canBeClimbed())
                        .setIsIce(blk.getId() == BlockID.ICE)
                        .setIsCobweb(blk.getId() == BlockID.COBWEB)
                        //no sweet berry in Nukkit
                        .setIsSweetBerry(false)
                        .build())
                .build();
    }

    @EventHandler
    public void onPlayerMove(PlayerMoveEvent ev) {
        val player = ev.getPlayer();
        val data = this.data.get(player);
        if (data != null) {
            data.getQueue().add(getTick(),
                    PlayerWrappers.WildcardReportData.newBuilder().setEffectData(
                            getEffectData(player)
                    ).build()
            );
            data.getQueue().add(getTick(),
                    PlayerWrappers.WildcardReportData.newBuilder().setMoveData(
                            getMovementData(player, ev.getTo(), false)
                    ).build()
            );
        }
    }

    @EventHandler
    public void onPlayerTeleport(EntityTeleportEvent ev) {
        if (ev.getEntity() instanceof Player) {
            val player = (Player) ev.getEntity();
            val data = this.data.get(player);
            if (data != null) {
                data.getQueue().add(getTick(),
                        PlayerWrappers.WildcardReportData.newBuilder().setMoveData(
                                getMovementData(player, ev.getTo(), true)
                        ).build()
                );
            }
        }
    }

    @EventHandler
    public void onPlayerToggleSprint(PlayerToggleSprintEvent ev) {
        val player = ev.getPlayer();
        val data = this.data.get(player);
        if (data != null) {
            data.getQueue().add(getTick(),
                    PlayerWrappers.WildcardReportData.newBuilder().setActionData(
                            getActionData(player,
                                    ev.isSprinting() ?
                                            PlayerOuterClass.PlayerAction.StartSprint :
                                            PlayerOuterClass.PlayerAction.StopSprint
                            )
                    ).build()
            );
        }
    }

    @EventHandler
    public void onPlayerToggleSneak(PlayerToggleSneakEvent ev) {
        val player = ev.getPlayer();
        val data = this.data.get(player);
        if (data != null) {
            data.getQueue().add(getTick(),
                    PlayerWrappers.WildcardReportData.newBuilder().setActionData(
                            getActionData(player,
                                    ev.isSneaking() ?
                                            PlayerOuterClass.PlayerAction.StartSneak :
                                            PlayerOuterClass.PlayerAction.StopSneak
                            )
                    ).build()
            );
        }
    }

    @EventHandler
    public void onPlayerJump(PlayerJumpEvent ev) {
        val player = ev.getPlayer();
        val data = this.data.get(player);
        if (data != null) {
            data.getQueue().add(getTick(),
                    PlayerWrappers.WildcardReportData.newBuilder().setActionData(
                            getActionData(player, PlayerOuterClass.PlayerAction.Jump)
                    ).build()
            );
        }
    }

    @EventHandler
    public void onPlayerDeath(PlayerDeathEvent ev) {
        val player = ev.getEntity();
        val data = this.data.get(player);
        if (data != null) {
            data.getQueue().add(getTick(),
                    PlayerWrappers.WildcardReportData.newBuilder().setLifeData(
                            getLifeData(false)
                    ).build()
            );
        }
    }

    @EventHandler
    public void onPlayerRespawn(PlayerRespawnEvent ev) {
        val player = ev.getPlayer();
        val data = this.data.get(player);
        if (data != null) {
            data.getQueue().add(getTick(),
                    PlayerWrappers.WildcardReportData.newBuilder().setLifeData(
                            getLifeData(true)
                    ).build()
            );
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

    private PlayerOuterClass.EntityPositionData getPositionData(Player player, Vector3 newPos) {
        var bb = player.getBoundingBox().clone();
        val delta = newPos.subtract(player.getPosition());
        bb = bb.getOffsetBoundingBox(delta.getX(), delta.getY(), delta.getZ());
        bb.setMinY(bb.getMinY() - 0.75);
        var collision = Arrays.asList(player.getLevel().getCollisionBlocks(bb));
        return PlayerOuterClass.EntityPositionData.newBuilder()
                .setPosition(toVec3f(player.getPosition()))
                .setDirection(toVec3f(player.getDirectionVector()))
                .setBoundingBox(toBB(player.getBoundingBox()))
                .setIsImmobile(player.isImmobile())
                .setIsOnGround(player.isOnGround())
                .setAllowFlying(player.getAdventureSettings().get(AdventureSettings.Type.ALLOW_FLIGHT))
                .setIsFlying(player.getAdventureSettings().get(AdventureSettings.Type.FLYING))
                //TODO improve this
                .setHaveGravity(true)
                .setMovementSpeed(player.getMovementSpeed())
                //TODO improve this
                .setWouldCollideVertically(player.isCollidedVertically)
                //TODO .setBelowThatAffectMovement(player.block)
                .addAllCollidedBlocks(
                        collision.stream()
                                .map(this::toBlock)
                                .collect(Collectors.toList())
                )
                //TODO .addIntersectedBlocks()
                .build();
    }

    private int getTick() {
        return Server.getInstance().getTick();
    }

    private PlayerWrappers.PlayerEffectData getEffectData(Player player) {
        return PlayerWrappers.PlayerEffectData.newBuilder().addAllEffect(getEffects(player)).build();
    }

    private List<PrimitiveTypes.EffectFeature> getEffects(Player player) {
        val l = new ArrayList<PrimitiveTypes.EffectFeature>(16);
        internalEff(l, player, Effect.SPEED, (e) -> e.setIsSpeed(true));
        internalEff(l, player, Effect.HASTE, (e) -> e.setIsHaste(true));
        internalEff(l, player, Effect.SLOW_FALLING, (e) -> e.setIsSlowFalling(true));
        internalEff(l, player, Effect.LEVITATION, (e) -> e.setIsLevitation(true));
        internalEff(l, player, Effect.SLOWNESS, (e) -> e.setIsSlowness(true));
        internalEff(l, player, Effect.JUMP_BOOST, (e) -> e.setIsJumpBoost(true));
        return l;
    }

    private void internalEff(List<PrimitiveTypes.EffectFeature> list, Player player, int effectId, Consumer<PrimitiveTypes.EffectFeature.Builder> c) {
        val eff = player.getEffect(effectId);
        if (eff != null) {
            val ef = PrimitiveTypes.EffectFeature.newBuilder()
                    .setAmplifier(eff.getAmplifier());
            c.accept(ef);
            list.add(ef.build());
        }
    }
}
