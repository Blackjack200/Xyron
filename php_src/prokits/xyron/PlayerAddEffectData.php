<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: player_wrappers.proto

namespace prokits\xyron;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>xchange.PlayerAddEffectData</code>
 */
class PlayerAddEffectData extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string effect = 1;</code>
     */
    protected $effect = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $effect
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\PlayerWrappers::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>string effect = 1;</code>
     * @return string
     */
    public function getEffect()
    {
        return $this->effect;
    }

    /**
     * Generated from protobuf field <code>string effect = 1;</code>
     * @param string $var
     * @return $this
     */
    public function setEffect($var)
    {
        GPBUtil::checkString($var, True);
        $this->effect = $var;

        return $this;
    }

}

