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
     * Generated from protobuf enum <code>DEBUG = 0;</code>
     */
    const DEBUG = 0;
    /**
     * Generated from protobuf enum <code>AMBIGUOUS = 1;</code>
     */
    const AMBIGUOUS = 1;
    /**
     * Generated from protobuf enum <code>TRIGGER = 2;</code>
     */
    const TRIGGER = 2;

    private static $valueToName = [
        self::DEBUG => 'DEBUG',
        self::AMBIGUOUS => 'AMBIGUOUS',
        self::TRIGGER => 'TRIGGER',
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

