package com.blackjack200.xyron.nukkit;

import cn.nukkit.Player;
import cn.nukkit.block.Block;
import cn.nukkit.block.BlockID;
import cn.nukkit.block.BlockLiquid;
import cn.nukkit.event.entity.EntityDamageEvent;
import cn.nukkit.item.*;
import cn.nukkit.math.AxisAlignedBB;
import cn.nukkit.math.SimpleAxisAlignedBB;
import cn.nukkit.math.Vector3;
import cn.nukkit.potion.Effect;
import com.github.blackjack200.xyron.PlayerOuterClass;
import com.github.blackjack200.xyron.PlayerWrappers;
import com.github.blackjack200.xyron.PrimitiveTypes;
import com.google.common.collect.Lists;
import lombok.experimental.UtilityClass;
import lombok.val;
import lombok.var;

import java.util.ArrayList;
import java.util.List;
import java.util.Map;
import java.util.function.Consumer;

@UtilityClass
public class Convert {

    public PlayerOuterClass.DeviceOS deviceOS(int os) {
        var c = PlayerOuterClass.DeviceOS.forNumber(os);
        if (c == null) {
            c = PlayerOuterClass.DeviceOS.Android;
        }
        return c;
    }

    public PlayerOuterClass.GameMode gameMode(int gameMode) {
        switch (gameMode) {
            case Player.CREATIVE:
                return PlayerOuterClass.GameMode.Creative;
            case Player.ADVENTURE:
                return PlayerOuterClass.GameMode.Adventure;
            case Player.SPECTATOR:
                return PlayerOuterClass.GameMode.Spectator;
            case Player.SURVIVAL:
                return PlayerOuterClass.GameMode.Survival;
            default:
                throw new RuntimeException("invalid gamemode: " + gameMode);
        }
    }

    public PlayerOuterClass.DamageCause damageCause(EntityDamageEvent.DamageCause cause) {
        if (cause == EntityDamageEvent.DamageCause.CONTACT) {
            return PlayerOuterClass.DamageCause.Contact;
        }
        if (cause == EntityDamageEvent.DamageCause.ENTITY_ATTACK) {
            return PlayerOuterClass.DamageCause.EntityAttack;
        }
        if (cause == EntityDamageEvent.DamageCause.PROJECTILE) {
            return PlayerOuterClass.DamageCause.Projectile;
        }
        if (cause == EntityDamageEvent.DamageCause.SUFFOCATION) {
            return PlayerOuterClass.DamageCause.Suffocation;
        }
        if (cause == EntityDamageEvent.DamageCause.FALL) {
            return PlayerOuterClass.DamageCause.Fall;
        }
        if (cause == EntityDamageEvent.DamageCause.FIRE) {
            return PlayerOuterClass.DamageCause.Fire;
        }
        if (cause == EntityDamageEvent.DamageCause.FIRE_TICK) {
            return PlayerOuterClass.DamageCause.FireTick;
        }
        if (cause == EntityDamageEvent.DamageCause.LAVA) {
            return PlayerOuterClass.DamageCause.Lava;
        }
        if (cause == EntityDamageEvent.DamageCause.DROWNING) {
            return PlayerOuterClass.DamageCause.Drowning;
        }
        if (cause == EntityDamageEvent.DamageCause.BLOCK_EXPLOSION) {
            return PlayerOuterClass.DamageCause.BlockExplosion;
        }
        if (cause == EntityDamageEvent.DamageCause.ENTITY_EXPLOSION) {
            return PlayerOuterClass.DamageCause.EntityExplosion;
        }
        if (cause == EntityDamageEvent.DamageCause.VOID) {
            return PlayerOuterClass.DamageCause.Void;
        }
        if (cause == EntityDamageEvent.DamageCause.SUICIDE) {
            return PlayerOuterClass.DamageCause.Suicide;
        }
        if (cause == EntityDamageEvent.DamageCause.MAGIC) {
            return PlayerOuterClass.DamageCause.Magic;
        }
        if (cause == EntityDamageEvent.DamageCause.CUSTOM) {
            return PlayerOuterClass.DamageCause.Custom;
        }
        if (cause == EntityDamageEvent.DamageCause.LIGHTNING) {
            return PlayerOuterClass.DamageCause.Custom;
        }
        if (cause == EntityDamageEvent.DamageCause.HUNGER) {
            return PlayerOuterClass.DamageCause.Starvation;
        }
        return PlayerOuterClass.DamageCause.Custom;
    }

    public PlayerOuterClass.InputMode inputMode(int mode) {
        var c = PlayerOuterClass.InputMode.forNumber(mode);
        if (c == null) {
            c = PlayerOuterClass.InputMode.Touch;
        }
        return c;
    }

    public PrimitiveTypes.Vec3f vec3f(Vector3 vec3) {
        return PrimitiveTypes.Vec3f.newBuilder()
                .setX((float) vec3.getX())
                .setY((float) vec3.getY())
                .setZ((float) vec3.getZ())
                .build();
    }

    public Vector3 vec3f(PrimitiveTypes.Vec3f v3f) {
        return new Vector3(v3f.getX(), v3f.getY(), v3f.getZ());
    }

    public PrimitiveTypes.Vec3i vec3i(Vector3 vec3) {
        return PrimitiveTypes.Vec3i.newBuilder()
                .setX((int) vec3.getX())
                .setY((int) vec3.getY())
                .setZ((int) vec3.getZ())
                .build();
    }

    public Vector3 vec3i(PrimitiveTypes.Vec3i vec3) {
        return new Vector3(vec3.getX(), vec3.getY(), vec3.getZ());
    }

    public PrimitiveTypes.AxisAlignedBoundingBox boundingBox(AxisAlignedBB aabb) {
        return PrimitiveTypes.AxisAlignedBoundingBox.newBuilder()
                .setMin(vec3f(new Vector3(aabb.getMinX(), aabb.getMinY(), aabb.getMinZ())))
                .setMax(vec3f(new Vector3(aabb.getMaxX(), aabb.getMaxY(), aabb.getMaxZ())))
                .build();
    }

    public AxisAlignedBB boundingBox(PrimitiveTypes.AxisAlignedBoundingBox aabb) {
        return new SimpleAxisAlignedBB(vec3f(aabb.getMin()), vec3f(aabb.getMax()));
    }

    public PrimitiveTypes.BlockData block(Block blk) {
        return PrimitiveTypes.BlockData.newBuilder()
                .setPosition(vec3i(blk))
                .setFeature(PrimitiveTypes.BlockFeature.newBuilder()
                        .addAllCollisionBoxes(Lists.newArrayList(boundingBox(blk.getCollisionBoundingBox())))
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
                        .build()
                ).build();
    }

    public PrimitiveTypes.ItemData item(Item item) {
        return PrimitiveTypes.ItemData.newBuilder()
                //FIXME unfortunately Nukkit don't use this
                .setVanillaName(item.getName())
                .setCount(item.getCount())
                .setFeature(PrimitiveTypes.ItemFeature.newBuilder()
                        .setIsArmor(item.isArmor())
                        .setIsBlockPlanterItem(false)
                        .setIsDamageable(item instanceof ItemDurable)
                        .setIsFood(item instanceof ItemEdible)
                        .setIsThrowable(item instanceof ProjectileItem)
                        .setIsTool(item.isTool())
                        .setIsBow(item instanceof ItemBow)
                        .setIsCrossBow(item instanceof ItemCrossbow)
                        .setIsShield(item instanceof ItemShield)
                        .build()
                )
                .build();
    }

    public List<PrimitiveTypes.EffectFeature> effects(Map<Integer, Effect> effects) {
        val l = new ArrayList<PrimitiveTypes.EffectFeature>(16);
        internalEffect(l, effects, Effect.SPEED, (e) -> e.setIsSpeed(true));
        internalEffect(l, effects, Effect.HASTE, (e) -> e.setIsHaste(true));
        internalEffect(l, effects, Effect.SLOW_FALLING, (e) -> e.setIsSlowFalling(true));
        internalEffect(l, effects, Effect.LEVITATION, (e) -> e.setIsLevitation(true));
        internalEffect(l, effects, Effect.SLOWNESS, (e) -> e.setIsSlowness(true));
        internalEffect(l, effects, Effect.JUMP_BOOST, (e) -> e.setIsJumpBoost(true));
        return l;
    }

    private void internalEffect(List<PrimitiveTypes.EffectFeature> list, Map<Integer, Effect> effects, int effectId, Consumer<PrimitiveTypes.EffectFeature.Builder> c) {
        val eff = effects.getOrDefault(effectId, null);
        if (eff != null) {
            val ef = PrimitiveTypes.EffectFeature.newBuilder()
                    .setAmplifier(eff.getAmplifier());
            c.accept(ef);
            list.add(ef.build());
        }
    }

    public PlayerWrappers.WildcardReportData wildcard(PlayerWrappers.PlayerActionData data) {
        return PlayerWrappers.WildcardReportData.newBuilder().setActionData(data).build();
    }

    public PlayerWrappers.WildcardReportData wildcard(PlayerWrappers.PlayerMoveData data) {
        return PlayerWrappers.WildcardReportData.newBuilder().setMoveData(data).build();
    }

    public PlayerWrappers.WildcardReportData wildcard(PlayerWrappers.PlayerPlaceBlockData data) {
        return PlayerWrappers.WildcardReportData.newBuilder().setPlaceBlockData(data).build();
    }

    public PlayerWrappers.WildcardReportData wildcard(PlayerWrappers.PlayerBreakBlockData data) {
        return PlayerWrappers.WildcardReportData.newBuilder().setBreakBlockData(data).build();
    }

    public PlayerWrappers.WildcardReportData wildcard(PlayerWrappers.PlayerEatFoodData data) {
        return PlayerWrappers.WildcardReportData.newBuilder().setEatFoodData(data).build();
    }

    public PlayerWrappers.WildcardReportData wildcard(PlayerWrappers.PlayerAttackData data) {
        return PlayerWrappers.WildcardReportData.newBuilder().setAttackData(data).build();
    }

    public PlayerWrappers.WildcardReportData wildcard(PlayerWrappers.PlayerEffectData data) {
        return PlayerWrappers.WildcardReportData.newBuilder().setEffectData(data).build();
    }

    public PlayerWrappers.WildcardReportData wildcard(PlayerWrappers.PlayerGameModeData data) {
        return PlayerWrappers.WildcardReportData.newBuilder().setGameModeData(data).build();
    }

    public PlayerWrappers.WildcardReportData wildcard(PlayerWrappers.PlayerMotionData data) {
        return PlayerWrappers.WildcardReportData.newBuilder().setMotionData(data).build();
    }

    public PlayerWrappers.WildcardReportData wildcard(PlayerWrappers.PlayerInputModeData data) {
        return PlayerWrappers.WildcardReportData.newBuilder().setInputModeData(data).build();
    }

    public PlayerWrappers.WildcardReportData wildcard(PlayerWrappers.PlayerHeldItemChangeData data) {
        return PlayerWrappers.WildcardReportData.newBuilder().setHeldItemChangeData(data).build();
    }

    public PlayerWrappers.WildcardReportData wildcard(PlayerWrappers.PlayerLifeData data) {
        return PlayerWrappers.WildcardReportData.newBuilder().setLifeData(data).build();
    }
}
