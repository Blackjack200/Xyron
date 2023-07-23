<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: anticheat_types.proto

namespace prokits\xyron;

use UnexpectedValueException;

/**
 * Protobuf type <code>xchange.Judgement</code>
 */
class Judgement
{
    /**
     * Generated from protobuf enum <code>NONE = 0;</code>
     */
    const NONE = 0;
    /**
     * Generated from protobuf enum <code>AMBIGUOUS = 1;</code>
     */
    const AMBIGUOUS = 1;
    /**
     * Generated from protobuf enum <code>BAN = 2;</code>
     */
    const BAN = 2;

    private static $valueToName = [
        self::NONE => 'NONE',
        self::AMBIGUOUS => 'AMBIGUOUS',
        self::BAN => 'BAN',
    ];

    public static function name($value)
    {
        if (!isset(self::$valueToName[$value])) {
            throw new UnexpectedValueException(sprintf(
                    'Enum %s has no name defined for value %s', __CLASS__, $value));
        }
        return self::$valueToName[$value];
    }


    public static function value($name)
    {
        $const = __CLASS__ . '::' . strtoupper($name);
        if (!defined($const)) {
            throw new UnexpectedValueException(sprintf(
                    'Enum %s has no value defined for name %s', __CLASS__, $name));
        }
        return constant($const);
    }
}
