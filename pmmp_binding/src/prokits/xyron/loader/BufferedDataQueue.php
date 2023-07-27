<?php

namespace prokits\xyron\loader;

use prokits\xyron\WildcardReportData;

class BufferedDataQueue {
	/** @var WildcardReportData[][] */
	private array $data = [];

	public function add(int $tick, WildcardReportData $data) : void {
		if (!isset($this->data[$tick])) {
			$this->data[$tick] = [];
		}
		$this->data[$tick][] = $data;
	}

	public function flush(int $tick) : \Generator {
		foreach ($this->data as $tck => $data) {
			if ($tck <= $tick) {
				yield $tck => $data;
				unset($this->data[$tck]);
			}
		}
	}
}