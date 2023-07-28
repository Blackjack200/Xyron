<?php

namespace prokits\xyron\loader;

use Closure;
use Generator;
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

	public function flush(int $tick, Closure &$remove) : Generator {
		$tks = [];
		foreach ($this->data as $tck => $data) {
			if ($tck <= $tick) {
				yield $tck => $data;
				$tks[] = $tck;
			}
		}
		$remove = function() use ($tks) : void {
			foreach ($tks as $tck) {
				unset($this->data[$tck]);
			}
		};
	}
}