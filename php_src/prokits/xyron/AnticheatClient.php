<?php
// GENERATED CODE -- DO NOT EDIT!

namespace prokits\xyron;

/**
 */
class AnticheatClient extends \Grpc\BaseStub {

    /**
     * @param string $hostname hostname
     * @param array $opts channel options
     * @param \Grpc\Channel $channel (optional) re-use channel object
     */
    public function __construct($hostname, $opts, $channel = null) {
        parent::__construct($hostname, $opts, $channel);
    }

    /**
     * @param \prokits\xyron\AddPlayerRequest $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function AddPlayer(\prokits\xyron\AddPlayerRequest $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/xchange.Anticheat/AddPlayer',
        $argument,
        ['\prokits\xyron\PlayerReceipt', 'decode'],
        $metadata, $options);
    }

    /**
     * @param \prokits\xyron\PlayerReceipt $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function RemovePlayer(\prokits\xyron\PlayerReceipt $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/xchange.Anticheat/RemovePlayer',
        $argument,
        ['\Google\Protobuf\GPBEmpty', 'decode'],
        $metadata, $options);
    }

    /**
     * @param \prokits\xyron\PlayerReport $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function Report(\prokits\xyron\PlayerReport $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/xchange.Anticheat/Report',
        $argument,
        ['\prokits\xyron\ReportResponse', 'decode'],
        $metadata, $options);
    }

}
