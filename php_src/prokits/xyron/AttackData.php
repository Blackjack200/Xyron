<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: player.proto

namespace prokits\xyron;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>xchange.AttackData</code>
 */
class AttackData extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>.xchange.DamageCause cause = 1;</code>
     */
    protected $cause = 0;
    /**
     * Generated from protobuf field <code>.xchange.EntityPositionData attacker = 2;</code>
     */
    protected $attacker = null;
    /**
     * Generated from protobuf field <code>.xchange.EntityPositionData target = 3;</code>
     */
    protected $target = null;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type int $cause
     *     @type \prokits\xyron\EntityPositionData $attacker
     *     @type \prokits\xyron\EntityPositionData $target
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Player::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>.xchange.DamageCause cause = 1;</code>
     * @return int
     */
    public function getCause()
    {
        return $this->cause;
    }

    /**
     * Generated from protobuf field <code>.xchange.DamageCause cause = 1;</code>
     * @param int $var
     * @return $this
     */
    public function setCause($var)
    {
        GPBUtil::checkEnum($var, \prokits\xyron\DamageCause::class);
        $this->cause = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>.xchange.EntityPositionData attacker = 2;</code>
     * @return \prokits\xyron\EntityPositionData|null
     */
    public function getAttacker()
    {
        return $this->attacker;
    }

    public function hasAttacker()
    {
        return isset($this->attacker);
    }

    public function clearAttacker()
    {
        unset($this->attacker);
    }

    /**
     * Generated from protobuf field <code>.xchange.EntityPositionData attacker = 2;</code>
     * @param \prokits\xyron\EntityPositionData $var
     * @return $this
     */
    public function setAttacker($var)
    {
        GPBUtil::checkMessage($var, \prokits\xyron\EntityPositionData::class);
        $this->attacker = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>.xchange.EntityPositionData target = 3;</code>
     * @return \prokits\xyron\EntityPositionData|null
     */
    public function getTarget()
    {
        return $this->target;
    }

    public function hasTarget()
    {
        return isset($this->target);
    }

    public function clearTarget()
    {
        unset($this->target);
    }

    /**
     * Generated from protobuf field <code>.xchange.EntityPositionData target = 3;</code>
     * @param \prokits\xyron\EntityPositionData $var
     * @return $this
     */
    public function setTarget($var)
    {
        GPBUtil::checkMessage($var, \prokits\xyron\EntityPositionData::class);
        $this->target = $var;

        return $this;
    }

}

