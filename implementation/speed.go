package implementation

import (
	"fmt"
	"github.com/blackjack200/xyron/anticheat"
	"github.com/blackjack200/xyron/xyron"
	"github.com/go-gl/mathgl/mgl64"
	"math"
)

type Speed struct {
	*anticheat.Evaluator
	PredictionLatitude float64
	UnstableRate       float64
}

var _ anticheat.MoveDataHandler = &Speed{}

func init() {
	oldA := Available
	Available = func() []any {
		return append(oldA(), &Speed{
			anticheat.NewEvaluator(80, 0.75, 0.96),
			0.05,
			0.997,
		})
	}
}

func (g *Speed) HandleMoveData(p *anticheat.InternalPlayer, data *xyron.PlayerMoveData) *xyron.JudgementData {
	if math.Abs(mgl64.Clamp(p.Motion.Current().Get().Y(), -epsilon, epsilon)) > epsilon {
		return nil
	}
	tickSinceTeleport := p.CurrentTimestamp() - p.Teleport.Current().Timestamp()
	if tickSinceTeleport < 15 {
		return nil
	}
	if p.Location.Previous() == nil ||
		//TODO better high jump support
		p.InAirTick < 15 ||
		p.IntersectedLiquid.Current().Get() ||
		(!p.Flying.Current().Get() && p.CurrentTimestamp()-p.Flying.Current().Timestamp() < 20) ||
		p.OnGround.Current().Get() {
		return nil
	}
	newOnGround, _, _, _, _, _ := p.CheckGroundState(data.NewPosition)
	//sometimes when player land and death, false positives appear
	if newOnGround {
		return nil
	}
	if p.Location.Current().IsFlying {
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
	prevDeltaXZ := math.Hypot(prevDelta.X(), prevDelta.Z())
	measuredDeltaXZ := math.Hypot(measuredDelta.X(), measuredDelta.Z())

	if newPos.Sub(oldPos).Len() <= epsilon {
		return nil
	}

	sp := 0.02
	if p.Sprinting.Current().Get() {
		sp = 0.026
	}
	predictedDeltaXZ := prevDeltaXZ*0.91 + sp

	equalness := 1 - math.Min(measuredDeltaXZ/predictedDeltaXZ, predictedDeltaXZ/measuredDeltaXZ)
	if !p.Location.Current().IsFlying && !data.NewPosition.IsFlying {
		g.HandleMaxRate(equalness, g.PredictionLatitude, g.UnstableRate)
	}
	return &xyron.JudgementData{
		Type:      "Speed",
		Judgement: g.Evaluate(),
		Message:   fmt.Sprintf("p:%v pred-xz:%.5f xz:%.5f eq:%.5f", g.PossibilityString(), predictedDeltaXZ, prevDeltaXZ, equalness),
	}
}
