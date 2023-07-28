package implementation

import (
	"fmt"
	"github.com/blackjack200/xyron/anticheat"
	"github.com/blackjack200/xyron/xyron"
	"github.com/go-gl/mathgl/mgl64"
	"math"
	"strconv"
)

type MovementPrediction struct {
	*anticheat.Evaluator
	pred               *predictor
	PredictionLatitude float64
	UnstableRate       float64
}

var _ anticheat.MoveDataHandler = &MovementPrediction{}

func (g *MovementPrediction) HandleMoveData(p *anticheat.InternalPlayer, data *xyron.PlayerMoveData) *xyron.JudgementData {
	if math.Abs(mgl64.Clamp(p.Motion.Current().Get().Y(), -epsilon, epsilon)) > epsilon {
		return nil
	}
	tickSinceTeleport := p.CurrentTimestamp() - p.Teleport.Current().Timestamp()
	if tickSinceTeleport < 15 {
		return nil
	}
	if p.Location.Previous() == nil {
		return nil
	}
	motion := p.Motion.Current()
	if motion.Get().Len() > epsilon {
		maxTick := int64(10)
		//TODO improve big motion support
		if motion.Get().Len() >= 1.5 {
			maxTick = 35
		}
		tickSinceMotion := p.CurrentTimestamp() - motion.Timestamp()
		if tickSinceMotion < maxTick {
			return nil
		}
	}

	oldOldPos := toVec3(p.Location.Previous().Position)
	oldPos := toVec3(p.Location.Current().Position)
	newPos := toVec3(data.NewPosition.Position)

	prevDelta := oldPos.Sub(oldOldPos)
	measuredDelta := newPos.Sub(oldPos)
	predictedDelta := g.pred.predictNextTickDeltaMovement(p, prevDelta)

	calcVec3Sim := func(a, b mgl64.Vec3) float64 {
		return b.Sub(a).Len()
	}

	roundVec3 := func(a mgl64.Vec3) mgl64.Vec3 {
		x, _ := strconv.ParseFloat(fmt.Sprintf("%.4f", a.X()), 64)
		y, _ := strconv.ParseFloat(fmt.Sprintf("%.4f", a.Y()), 64)
		z, _ := strconv.ParseFloat(fmt.Sprintf("%.4f", a.Z()), 64)
		return mgl64.Vec3{
			x, y, z,
		}
	}

	sim := calcVec3Sim(measuredDelta, predictedDelta)
	/*
		if !p.Location.Current().IsFlying {
			g.HandleUnstableRate(sim, g.PredictionLatitude, g.UnstableRate)
		}
	*/
	return &xyron.JudgementData{
		Type:      "MovementPrediction",
		Judgement: g.Evaluate(),
		Message:   fmt.Sprintf("p:%v pred-d:%v d:%v eq:%.4f", g.PossibilityString(), roundVec3(predictedDelta), roundVec3(measuredDelta), sim),
	}
}
