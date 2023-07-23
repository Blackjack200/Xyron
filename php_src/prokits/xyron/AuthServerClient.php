<?php
// GENERATED CODE -- DO NOT EDIT!

namespace prokits\xyron;

/**
 */
class AuthServerClient extends \Grpc\BaseStub {

    /**
     * @param string $hostname hostname
     * @param array $opts channel options
     * @param \Grpc\Channel $channel (optional) re-use channel object
     */
    public function __construct($hostname, $opts, $channel = null) {
        parent::__construct($hostname, $opts, $channel);
    }

    /**
     * @param \prokits\xyron\ClientInfo $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function Login(\prokits\xyron\ClientInfo $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/xchange.AuthServer/Login',
        $argument,
        ['\prokits\xyron\ClientReceipt', 'decode'],
        $metadata, $options);
    }

    /**
     * @param \prokits\xyron\ClientInfo $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function Logout(\prokits\xyron\ClientInfo $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/xchange.AuthServer/Logout',
        $argument,
        ['\Google\Protobuf\GPBEmpty', 'decode'],
        $metadata, $options);
    }

}
