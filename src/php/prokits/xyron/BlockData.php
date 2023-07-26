<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: primitive_types.proto

namespace prokits\xyron;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>xchange.BlockData</code>
 */
class BlockData extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>.xchange.BlockFeature feature = 1;</code>
     */
    protected $feature = null;
    /**
     * Generated from protobuf field <code>.xchange.Vec3i position = 2;</code>
     */
    protected $position = null;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type \prokits\xyron\BlockFeature $feature
     *     @type \prokits\xyron\Vec3i $position
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\PrimitiveTypes::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>.xchange.BlockFeature feature = 1;</code>
     * @return \prokits\xyron\BlockFeature|null
     */
    public function getFeature()
    {
        return $this->feature;
    }

    public function hasFeature()
    {
        return isset($this->feature);
    }

    public function clearFeature()
    {
        unset($this->feature);
    }

    /**
     * Generated from protobuf field <code>.xchange.BlockFeature feature = 1;</code>
     * @param \prokits\xyron\BlockFeature $var
     * @return $this
     */
    public function setFeature($var)
    {
        GPBUtil::checkMessage($var, \prokits\xyron\BlockFeature::class);
        $this->feature = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>.xchange.Vec3i position = 2;</code>
     * @return \prokits\xyron\Vec3i|null
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
     * Generated from protobuf field <code>.xchange.Vec3i position = 2;</code>
     * @param \prokits\xyron\Vec3i $var
     * @return $this
     */
    public function setPosition($var)
    {
        GPBUtil::checkMessage($var, \prokits\xyron\Vec3i::class);
        $this->position = $var;

        return $this;
    }

}

