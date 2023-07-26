<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: player_wrappers.proto

namespace prokits\xyron;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>xchange.WildcardReportData</code>
 */
class WildcardReportData extends \Google\Protobuf\Internal\Message
{
    protected $data;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type \prokits\xyron\PlayerActionData $action_data
     *     @type \prokits\xyron\PlayerMoveData $move_data
     *     @type \prokits\xyron\PlayerPlaceBlockData $place_block_data
     *     @type \prokits\xyron\PlayerBreakBlockData $break_block_data
     *     @type \prokits\xyron\PlayerEatFoodData $eat_food_data
     *     @type \prokits\xyron\PlayerAttackData $attack_data
     *     @type \prokits\xyron\PlayerEffectData $effect_data
     *     @type \prokits\xyron\PlayerGameModeData $game_mode_data
     *     @type \prokits\xyron\PlayerMotionData $motion_data
     *     @type \prokits\xyron\PlayerInputModeData $input_mode_data
     *     @type \prokits\xyron\PlayerHeldItemChangeData $held_item_change_data
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\PlayerWrappers::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>.xchange.PlayerActionData action_data = 1;</code>
     * @return \prokits\xyron\PlayerActionData|null
     */
    public function getActionData()
    {
        return $this->readOneof(1);
    }

    public function hasActionData()
    {
        return $this->hasOneof(1);
    }

    /**
     * Generated from protobuf field <code>.xchange.PlayerActionData action_data = 1;</code>
     * @param \prokits\xyron\PlayerActionData $var
     * @return $this
     */
    public function setActionData($var)
    {
        GPBUtil::checkMessage($var, \prokits\xyron\PlayerActionData::class);
        $this->writeOneof(1, $var);

        return $this;
    }

    /**
     * Generated from protobuf field <code>.xchange.PlayerMoveData move_data = 2;</code>
     * @return \prokits\xyron\PlayerMoveData|null
     */
    public function getMoveData()
    {
        return $this->readOneof(2);
    }

    public function hasMoveData()
    {
        return $this->hasOneof(2);
    }

    /**
     * Generated from protobuf field <code>.xchange.PlayerMoveData move_data = 2;</code>
     * @param \prokits\xyron\PlayerMoveData $var
     * @return $this
     */
    public function setMoveData($var)
    {
        GPBUtil::checkMessage($var, \prokits\xyron\PlayerMoveData::class);
        $this->writeOneof(2, $var);

        return $this;
    }

    /**
     * Generated from protobuf field <code>.xchange.PlayerPlaceBlockData place_block_data = 3;</code>
     * @return \prokits\xyron\PlayerPlaceBlockData|null
     */
    public function getPlaceBlockData()
    {
        return $this->readOneof(3);
    }

    public function hasPlaceBlockData()
    {
        return $this->hasOneof(3);
    }

    /**
     * Generated from protobuf field <code>.xchange.PlayerPlaceBlockData place_block_data = 3;</code>
     * @param \prokits\xyron\PlayerPlaceBlockData $var
     * @return $this
     */
    public function setPlaceBlockData($var)
    {
        GPBUtil::checkMessage($var, \prokits\xyron\PlayerPlaceBlockData::class);
        $this->writeOneof(3, $var);

        return $this;
    }

    /**
     * Generated from protobuf field <code>.xchange.PlayerBreakBlockData break_block_data = 4;</code>
     * @return \prokits\xyron\PlayerBreakBlockData|null
     */
    public function getBreakBlockData()
    {
        return $this->readOneof(4);
    }

    public function hasBreakBlockData()
    {
        return $this->hasOneof(4);
    }

    /**
     * Generated from protobuf field <code>.xchange.PlayerBreakBlockData break_block_data = 4;</code>
     * @param \prokits\xyron\PlayerBreakBlockData $var
     * @return $this
     */
    public function setBreakBlockData($var)
    {
        GPBUtil::checkMessage($var, \prokits\xyron\PlayerBreakBlockData::class);
        $this->writeOneof(4, $var);

        return $this;
    }

    /**
     * Generated from protobuf field <code>.xchange.PlayerEatFoodData eat_food_data = 5;</code>
     * @return \prokits\xyron\PlayerEatFoodData|null
     */
    public function getEatFoodData()
    {
        return $this->readOneof(5);
    }

    public function hasEatFoodData()
    {
        return $this->hasOneof(5);
    }

    /**
     * Generated from protobuf field <code>.xchange.PlayerEatFoodData eat_food_data = 5;</code>
     * @param \prokits\xyron\PlayerEatFoodData $var
     * @return $this
     */
    public function setEatFoodData($var)
    {
        GPBUtil::checkMessage($var, \prokits\xyron\PlayerEatFoodData::class);
        $this->writeOneof(5, $var);

        return $this;
    }

    /**
     * Generated from protobuf field <code>.xchange.PlayerAttackData attack_data = 6;</code>
     * @return \prokits\xyron\PlayerAttackData|null
     */
    public function getAttackData()
    {
        return $this->readOneof(6);
    }

    public function hasAttackData()
    {
        return $this->hasOneof(6);
    }

    /**
     * Generated from protobuf field <code>.xchange.PlayerAttackData attack_data = 6;</code>
     * @param \prokits\xyron\PlayerAttackData $var
     * @return $this
     */
    public function setAttackData($var)
    {
        GPBUtil::checkMessage($var, \prokits\xyron\PlayerAttackData::class);
        $this->writeOneof(6, $var);

        return $this;
    }

    /**
     * Generated from protobuf field <code>.xchange.PlayerEffectData effect_data = 7;</code>
     * @return \prokits\xyron\PlayerEffectData|null
     */
    public function getEffectData()
    {
        return $this->readOneof(7);
    }

    public function hasEffectData()
    {
        return $this->hasOneof(7);
    }

    /**
     * Generated from protobuf field <code>.xchange.PlayerEffectData effect_data = 7;</code>
     * @param \prokits\xyron\PlayerEffectData $var
     * @return $this
     */
    public function setEffectData($var)
    {
        GPBUtil::checkMessage($var, \prokits\xyron\PlayerEffectData::class);
        $this->writeOneof(7, $var);

        return $this;
    }

    /**
     * Generated from protobuf field <code>.xchange.PlayerGameModeData game_mode_data = 9;</code>
     * @return \prokits\xyron\PlayerGameModeData|null
     */
    public function getGameModeData()
    {
        return $this->readOneof(9);
    }

    public function hasGameModeData()
    {
        return $this->hasOneof(9);
    }

    /**
     * Generated from protobuf field <code>.xchange.PlayerGameModeData game_mode_data = 9;</code>
     * @param \prokits\xyron\PlayerGameModeData $var
     * @return $this
     */
    public function setGameModeData($var)
    {
        GPBUtil::checkMessage($var, \prokits\xyron\PlayerGameModeData::class);
        $this->writeOneof(9, $var);

        return $this;
    }

    /**
     * Generated from protobuf field <code>.xchange.PlayerMotionData motion_data = 10;</code>
     * @return \prokits\xyron\PlayerMotionData|null
     */
    public function getMotionData()
    {
        return $this->readOneof(10);
    }

    public function hasMotionData()
    {
        return $this->hasOneof(10);
    }

    /**
     * Generated from protobuf field <code>.xchange.PlayerMotionData motion_data = 10;</code>
     * @param \prokits\xyron\PlayerMotionData $var
     * @return $this
     */
    public function setMotionData($var)
    {
        GPBUtil::checkMessage($var, \prokits\xyron\PlayerMotionData::class);
        $this->writeOneof(10, $var);

        return $this;
    }

    /**
     * Generated from protobuf field <code>.xchange.PlayerInputModeData input_mode_data = 11;</code>
     * @return \prokits\xyron\PlayerInputModeData|null
     */
    public function getInputModeData()
    {
        return $this->readOneof(11);
    }

    public function hasInputModeData()
    {
        return $this->hasOneof(11);
    }

    /**
     * Generated from protobuf field <code>.xchange.PlayerInputModeData input_mode_data = 11;</code>
     * @param \prokits\xyron\PlayerInputModeData $var
     * @return $this
     */
    public function setInputModeData($var)
    {
        GPBUtil::checkMessage($var, \prokits\xyron\PlayerInputModeData::class);
        $this->writeOneof(11, $var);

        return $this;
    }

    /**
     * Generated from protobuf field <code>.xchange.PlayerHeldItemChangeData held_item_change_data = 12;</code>
     * @return \prokits\xyron\PlayerHeldItemChangeData|null
     */
    public function getHeldItemChangeData()
    {
        return $this->readOneof(12);
    }

    public function hasHeldItemChangeData()
    {
        return $this->hasOneof(12);
    }

    /**
     * Generated from protobuf field <code>.xchange.PlayerHeldItemChangeData held_item_change_data = 12;</code>
     * @param \prokits\xyron\PlayerHeldItemChangeData $var
     * @return $this
     */
    public function setHeldItemChangeData($var)
    {
        GPBUtil::checkMessage($var, \prokits\xyron\PlayerHeldItemChangeData::class);
        $this->writeOneof(12, $var);

        return $this;
    }

    /**
     * @return string
     */
    public function getData()
    {
        return $this->whichOneof("data");
    }

}

