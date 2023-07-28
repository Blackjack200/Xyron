package anticheat

import (
	"fmt"
	"github.com/blackjack200/xyron/xyron"
	"math"
)

type ViolationBuffer struct {
	buf float64
	Max float64
}

func NewViolationBuffer(max float64) *ViolationBuffer {
	return &ViolationBuffer{Max: max}
}

func (b *ViolationBuffer) Possibility() float64 {
	return math.Min(1, b.buf/b.Max)
}

func (b *ViolationBuffer) PossibilityString() string {
	return fmt.Sprintf("%.3f%%", math.Min(1, b.buf/b.Max)*100)
}

func (b *ViolationBuffer) HandleMax(measured, expectedMax float64) {
	if measured > expectedMax {
		b.buf++
	}
}

func (b *ViolationBuffer) HandleMaxRate(measured, expectedMax, rate float64) {
	if measured > expectedMax {
		b.buf++
	} else {
		b.buf *= rate
	}
}

func (b *ViolationBuffer) Add() {
	b.buf++
}

func (b *ViolationBuffer) HandleUnstable(measured, expectedMax float64) {
	if measured > expectedMax {
		b.buf++
	} else {
		b.buf = 0
	}
}

func (b *ViolationBuffer) HandleUnstableRate(measured, expectedMax, rate float64) {
	if measured > expectedMax {
		b.buf++
	} else {
		b.buf *= rate
	}
}

func (b *ViolationBuffer) HandleRelative(measured, expected, latitude float64) {
	if math.Abs(measured-expected) > latitude {
		b.buf++
	}
}

func (b *ViolationBuffer) HandleRelativeUnstable(measured, expected, latitude float64) {
	b.HandleRelativeUnstableRate(measured, expected, latitude, 0)
}

func (b *ViolationBuffer) HandleRelativeUnstableRate(measured, expected, latitude, rate float64) {
	if math.Abs(measured-expected) > latitude {
		b.buf++
	} else {
		b.buf *= rate
	}
}

type Evaluator struct {
	*ViolationBuffer
	MinValidPossibility     float64
	MaxPossibilityAmbiguous float64
}

func NewEvaluator(max float64, minValidPossibility float64, maxPossibilityAmbiguous float64) *Evaluator {
	return &Evaluator{
		ViolationBuffer:         NewViolationBuffer(max),
		MinValidPossibility:     minValidPossibility,
		MaxPossibilityAmbiguous: maxPossibilityAmbiguous,
	}
}

func (e *Evaluator) Evaluate() xyron.Judgement {
	possibility := e.Possibility()
	if possibility <= e.MinValidPossibility {
		return xyron.Judgement_DEBUG
	}
	if possibility <= e.MaxPossibilityAmbiguous {
		return xyron.Judgement_AMBIGUOUS
	}
	return xyron.Judgement_TRIGGER
}
