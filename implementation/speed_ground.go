package implementation

import (
	"fmt"
	"github.com/blackjack200/xyron/anticheat"
	"github.com/blackjack200/xyron/xyron"
	"github.com/go-gl/mathgl/mgl64"
	"math"
)

type SpeedGround struct {
	*anticheat.Evaluator
	PredictionLatitude float64
	UnstableRate       float64
}

var _ anticheat.MoveDataHandler = &SpeedGround{}

func init() {
	register(func() any {
		return &SpeedGround{
			anticheat.NewEvaluator(80, 0.75, 0.96),
			0.05,
			0.997,
		}
	})
}

func (g *SpeedGround) HandleMoveData(p *anticheat.InternalPlayer, data *xyron.PlayerMoveData) *xyron.JudgementData {
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

	if isZero(futurePos.Sub(pos).Len()) {
		return nil
	}

	slipperness := getSlipperiness(p.Location.Current().BelowThatAffectMovement.Feature)

	//TODO movement direction: https://www.mcpk.wiki/wiki/Horizontal_Movement_Formulas/zh
	// currently we use maximum value as possible
	movementFactor := 0.98
	if p.Sprinting.Current().Get() {
		movementFactor = 1.3
	}
	if p.Sneaking.Current().Get() {
		movementFactor = 0.3 * 0.98 * math.Sqrt(2)
	}

	effectsFactor := 1.0

	if e, ok := p.Effect(func(f *xyron.EffectFeature) bool {
		return f.IsSpeed
	}); ok {
		effectsFactor *= 1 + 0.2*float64(e.Amplifier)
	}
	if e, ok := p.Effect(func(f *xyron.EffectFeature) bool {
		return f.IsSlowness
	}); ok {
		effectsFactor *= 1 - 0.15*float64(e.Amplifier)
	}

	slippernessFactor := math.Pow(0.6/slipperness, 3)
	predictedMaxDX := delta.X()*slipperness*0.91 + 0.1*movementFactor*effectsFactor*slippernessFactor
	predictedMaxDZ := delta.Z()*slipperness*0.91 + 0.1*movementFactor*effectsFactor*slippernessFactor
	pred := mgl64.Vec2{predictedMaxDX, predictedMaxDZ}.Len()
	measured := mgl64.Vec2{deltaFuture.X(), deltaFuture.Z()}.Len()

	if !p.Location.Current().IsFlying &&
		!data.NewPosition.IsFlying &&
		!p.Location.Current().AllowFlying {
		g.HandleRelativeUnstableRate(measured, pred, g.PredictionLatitude, g.UnstableRate)
	}

	equalness := math.Abs(measured - pred)
	return &xyron.JudgementData{
		Type:      "SpeedGround",
		Judgement: g.Evaluate(),
		Message:   fmt.Sprintf("p:%v pred-xz:%.5f xz:%.5f delta:%.5f", g.PossibilityString(), measured, pred, equalness),
	}
}

func getSlipperiness(feature *xyron.BlockFeature) float64 {
	sliperness := 0.6
	if feature.IsAir {
		sliperness = 1.0
	}
	if feature.IsSlime {
		sliperness = 0.8
	}
	if feature.IsIce {
		sliperness = 0.98
	}
	return sliperness
}
