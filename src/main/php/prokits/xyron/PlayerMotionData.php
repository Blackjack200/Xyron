<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: player_wrappers.proto

namespace prokits\xyron;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>xchange.PlayerMotionData</code>
 */
class PlayerMotionData extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>.xchange.EntityPositionData position = 1;</code>
     */
    protected $position = null;
    /**
     * Generated from protobuf field <code>.xchange.Vec3f motion = 2;</code>
     */
    protected $motion = null;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type \prokits\xyron\EntityPositionData $position
     *     @type \prokits\xyron\Vec3f $motion
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\PlayerWrappers::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>.xchange.EntityPositionData position = 1;</code>
     * @return \prokits\xyron\EntityPositionData|null
     */
    public function getPosition()
    {
        return $this->position;
    }

    public function hasPosition()
    {
        return isset($this->position);
    }

    public function clearPosition()
    {
        unset($this->position);
    }

    /**
     * Generated from protobuf field <code>.xchange.EntityPositionData position = 1;</code>
     * @param \prokits\xyron\EntityPositionData $var
     * @return $this
     */
    public function setPosition($var)
    {
        GPBUtil::checkMessage($var, \prokits\xyron\EntityPositionData::class);
        $this->position = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>.xchange.Vec3f motion = 2;</code>
     * @return \prokits\xyron\Vec3f|null
     */
    public function getMotion()
    {
        return $this->motion;
    }

    public function hasMotion()
    {
        return isset($this->motion);
    }

    public function clearMotion()
    {
        unset($this->motion);
    }

    /**
     * Generated from protobuf field <code>.xchange.Vec3f motion = 2;</code>
     * @param \prokits\xyron\Vec3f $var
     * @return $this
     */
    public function setMotion($var)
    {
        GPBUtil::checkMessage($var, \prokits\xyron\Vec3f::class);
        $this->motion = $var;

        return $this;
    }

}

