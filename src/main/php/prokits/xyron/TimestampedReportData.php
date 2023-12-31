<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: xchange.proto

namespace prokits\xyron;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>xchange.TimestampedReportData</code>
 */
class TimestampedReportData extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>repeated .xchange.WildcardReportData data = 1;</code>
     */
    private $data;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type array<\prokits\xyron\WildcardReportData>|\Google\Protobuf\Internal\RepeatedField $data
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Xchange::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>repeated .xchange.WildcardReportData data = 1;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getData()
    {
        return $this->data;
    }

    /**
     * Generated from protobuf field <code>repeated .xchange.WildcardReportData data = 1;</code>
     * @param array<\prokits\xyron\WildcardReportData>|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setData($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \prokits\xyron\WildcardReportData::class);
        $this->data = $arr;

        return $this;
    }

}

