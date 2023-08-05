package implementation

import (
	"fmt"
	"github.com/blackjack200/xyron/anticheat"
	"github.com/blackjack200/xyron/xyron"
	"math"
)

type SpeedAir struct {
	*anticheat.Evaluator
	PredictionLatitude float64
	UnstableRate       float64
}

var _ anticheat.MoveDataHandler = &SpeedAir{}

func init() {
	register(func() any {
		return &SpeedAir{
			anticheat.NewEvaluator(80, 0.75, 0.96),
			0.05,
			0.997,
		}
	})
}

func (g *SpeedAir) HandleMoveData(p *anticheat.InternalPlayer, data *xyron.PlayerMoveData) *xyron.JudgementData {
	if !isPlayerFreeFalling(p, data.NewPosition) {
		return nil
	}
	if p.Location.Previous() == nil {
		return nil
	}
	oldPos := toVec3(p.Location.Previous().Position)
	pos := toVec3(p.Location.Current().Position)
	delta := pos.Sub(oldPos)

	futurePos := toVec3(data.NewPosition.Position)
	deltaFuture := futurePos.Sub(oldPos)

	deltaXZ := math.Hypot(delta.X(), delta.Z())
	measuredFutureDeltaXZ := math.Hypot(deltaFuture.X(), deltaFuture.Z())
	if isZero(futurePos.Sub(pos).Len()) {
		return nil
	}

	factor := 0.02
	if p.Sprinting.Current().Get() {
		factor = 0.026
	}
	predictedMaxDeltaXZ := deltaXZ*0.91 + factor

	if !p.Location.Current().IsFlying &&
		!data.NewPosition.IsFlying &&
		!p.Location.Current().AllowFlying {
		g.HandleRelativeUnstableRate(measuredFutureDeltaXZ, predictedMaxDeltaXZ, g.PredictionLatitude, g.UnstableRate)
	}

	equalness := math.Abs(measuredFutureDeltaXZ - predictedMaxDeltaXZ)
	return &xyron.JudgementData{
		Type:      "SpeedAir",
		Judgement: g.Evaluate(),
		Message:   fmt.Sprintf("p:%v pred-xz:%.5f xz:%.5f delta:%.5f", g.PossibilityString(), predictedMaxDeltaXZ, deltaXZ, equalness),
	}
}
