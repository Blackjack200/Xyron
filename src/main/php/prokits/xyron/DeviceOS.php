<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: player.proto

namespace prokits\xyron;

use UnexpectedValueException;

/**
 * Protobuf type <code>xchange.DeviceOS</code>
 */
class DeviceOS
{
    /**
     * Generated from protobuf enum <code>Android = 0;</code>
     */
    const Android = 0;
    /**
     * Generated from protobuf enum <code>IOS = 1;</code>
     */
    const IOS = 1;
    /**
     * Generated from protobuf enum <code>OSX = 2;</code>
     */
    const OSX = 2;
    /**
     * Generated from protobuf enum <code>AMAZON = 3;</code>
     */
    const AMAZON = 3;
    /**
     * Generated from protobuf enum <code>GEAR_VR = 4;</code>
     */
    const GEAR_VR = 4;
    /**
     * Generated from protobuf enum <code>HOLOLENS = 5;</code>
     */
    const HOLOLENS = 5;
    /**
     * Generated from protobuf enum <code>WINDOWS_10 = 6;</code>
     */
    const WINDOWS_10 = 6;
    /**
     * Generated from protobuf enum <code>WIN32 = 7;</code>
     */
    const WIN32 = 7;
    /**
     * Generated from protobuf enum <code>DEDICATED = 8;</code>
     */
    const DEDICATED = 8;
    /**
     * Generated from protobuf enum <code>TVOS = 9;</code>
     */
    const TVOS = 9;
    /**
     * Generated from protobuf enum <code>PLAYSTATION = 10;</code>
     */
    const PLAYSTATION = 10;
    /**
     * Generated from protobuf enum <code>NINTENDO = 11;</code>
     */
    const NINTENDO = 11;
    /**
     * Generated from protobuf enum <code>XBOX = 12;</code>
     */
    const XBOX = 12;
    /**
     * Generated from protobuf enum <code>WINDOWS_PHONE = 13;</code>
     */
    const WINDOWS_PHONE = 13;

    private static $valueToName = [
        self::Android => 'Android',
        self::IOS => 'IOS',
        self::OSX => 'OSX',
        self::AMAZON => 'AMAZON',
        self::GEAR_VR => 'GEAR_VR',
        self::HOLOLENS => 'HOLOLENS',
        self::WINDOWS_10 => 'WINDOWS_10',
        self::WIN32 => 'WIN32',
        self::DEDICATED => 'DEDICATED',
        self::TVOS => 'TVOS',
        self::PLAYSTATION => 'PLAYSTATION',
        self::NINTENDO => 'NINTENDO',
        self::XBOX => 'XBOX',
        self::WINDOWS_PHONE => 'WINDOWS_PHONE',
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

