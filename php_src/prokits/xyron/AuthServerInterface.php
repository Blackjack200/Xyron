<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: auth.proto

namespace prokits\xyron;

/**
 * Protobuf type <code>xchange.AuthServer</code>
 */
interface AuthServerInterface
{
    /**
     * Method <code>login</code>
     *
     * @param \prokits\xyron\ClientInfo $request
     * @return \prokits\xyron\ClientReceipt
     */
    public function login(\prokits\xyron\ClientInfo $request);

    /**
     * Method <code>logout</code>
     *
     * @param \prokits\xyron\ClientInfo $request
     * @return \Google\Protobuf\GPBEmpty
     */
    public function logout(\prokits\xyron\ClientInfo $request);

}
