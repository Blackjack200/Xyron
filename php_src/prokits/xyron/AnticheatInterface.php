<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: xchange.proto

namespace prokits\xyron;

/**
 * Protobuf type <code>xchange.Anticheat</code>
 */
interface AnticheatInterface
{
    /**
     * Method <code>addPlayer</code>
     *
     * @param \prokits\xyron\AddPlayerRequest $request
     * @return \prokits\xyron\PlayerReceipt
     */
    public function addPlayer(\prokits\xyron\AddPlayerRequest $request);

    /**
     * Method <code>removePlayer</code>
     *
     * @param \prokits\xyron\PlayerReceipt $request
     * @return \Google\Protobuf\GPBEmpty
     */
    public function removePlayer(\prokits\xyron\PlayerReceipt $request);

    /**
     * Method <code>report</code>
     *
     * @param \prokits\xyron\PlayerReport $request
     * @return \prokits\xyron\ReportResponse
     */
    public function report(\prokits\xyron\PlayerReport $request);

}

