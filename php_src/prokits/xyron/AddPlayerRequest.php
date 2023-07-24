<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: xchange.proto

namespace prokits\xyron;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>xchange.AddPlayerRequest</code>
 */
class AddPlayerRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>.xchange.Player player = 1;</code>
     */
    protected $player = null;
    /**
     *timestamp->report data
     *
     * Generated from protobuf field <code>map<int64, .xchange.TimestampedReportData> data = 2;</code>
     */
    private $data;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type \prokits\xyron\Player $player
     *     @type array|\Google\Protobuf\Internal\MapField $data
     *          timestamp->report data
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Xchange::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>.xchange.Player player = 1;</code>
     * @return \prokits\xyron\Player|null
     */
    public function getPlayer()
    {
        return $this->player;
    }

    public function hasPlayer()
    {
        return isset($this->player);
    }

    public function clearPlayer()
    {
        unset($this->player);
    }

    /**
     * Generated from protobuf field <code>.xchange.Player player = 1;</code>
     * @param \prokits\xyron\Player $var
     * @return $this
     */
    public function setPlayer($var)
    {
        GPBUtil::checkMessage($var, \prokits\xyron\Player::class);
        $this->player = $var;

        return $this;
    }

    /**
     *timestamp->report data
     *
     * Generated from protobuf field <code>map<int64, .xchange.TimestampedReportData> data = 2;</code>
     * @return \Google\Protobuf\Internal\MapField
     */
    public function getData()
    {
        return $this->data;
    }

    /**
     *timestamp->report data
     *
     * Generated from protobuf field <code>map<int64, .xchange.TimestampedReportData> data = 2;</code>
     * @param array|\Google\Protobuf\Internal\MapField $var
     * @return $this
     */
    public function setData($var)
    {
        $arr = GPBUtil::checkMapField($var, \Google\Protobuf\Internal\GPBType::INT64, \Google\Protobuf\Internal\GPBType::MESSAGE, \prokits\xyron\TimestampedReportData::class);
        $this->data = $arr;

        return $this;
    }

}

