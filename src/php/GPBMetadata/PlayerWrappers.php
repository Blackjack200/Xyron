<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: player_wrappers.proto

namespace GPBMetadata;

class PlayerWrappers
{
    public static $is_initialized = false;

    public static function initOnce() {
        $pool = \Google\Protobuf\Internal\DescriptorPool::getGeneratedPool();

        if (static::$is_initialized == true) {
          return;
        }
        \GPBMetadata\Player::initOnce();
        \GPBMetadata\PrimitiveTypes::initOnce();
        $pool->internalAddGeneratedFile(
            '
�
player_wrappers.protoxchangeprimitive_types.proto"�
WildcardReportData0
action_data (2.xchange.PlayerActionDataH ,
	move_data (2.xchange.PlayerMoveDataH 9
place_block_data (2.xchange.PlayerPlaceBlockDataH 9
break_block_data (2.xchange.PlayerBreakBlockDataH 3
eat_food_data (2.xchange.PlayerEatFoodDataH 0
attack_data (2.xchange.PlayerAttackDataH 0
effect_data (2.xchange.PlayerEffectDataH 5
game_mode_data	 (2.xchange.PlayerGameModeDataH 0
motion_data
 (2.xchange.PlayerMotionDataH 7
input_mode_data (2.xchange.PlayerInputModeDataH B
held_item_change_data (2!.xchange.PlayerHeldItemChangeDataH B
data"h
PlayerActionData-
position (2.xchange.EntityPositionData%
action (2.xchange.PlayerAction"T
PlayerMoveData0
newPosition (2.xchange.EntityPositionData
teleport ("n
PlayerPlaceBlockData-
position (2.xchange.EntityPositionData\'
placedBlock (2.xchange.BlockData"n
PlayerBreakBlockData-
position (2.xchange.EntityPositionData\'
brokenBlock (2.xchange.BlockData";
PlayerEatFoodData&
status (2.xchange.ConsumeStatus"5
PlayerAttackData!
data (2.xchange.AttackData":
PlayerEffectData&
effect (2.xchange.EffectFeature"9
PlayerGameModeData#
gameMode (2.xchange.GameMode"2
PlayerMotionData
motion (2.xchange.Vec3f"<
PlayerInputModeData%
	inputMode (2.xchange.InputMode";
PlayerHeldItemChangeData
item (2.xchange.ItemDataB=
com.github.blackjack200.xyronZxyron/��prokits\\xyron�bproto3'
        , true);

        static::$is_initialized = true;
    }
}

