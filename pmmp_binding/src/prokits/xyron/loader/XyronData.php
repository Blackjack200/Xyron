<?php

namespace prokits\xyron\loader;

use prokits\xyron\PlayerReceipt;

class XyronData {
	public function __construct(
		public PlayerReceipt     $receipt,
		public BufferedDataQueue $queue,
	) {
	}

	public function getQueue() : BufferedDataQueue { return $this->queue; }
}