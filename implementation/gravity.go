package implementation

import (
	"fmt"
	"github.com/blackjack200/xyron/anticheat"
	"github.com/blackjack200/xyron/xyron"
	"math"
)

type Gravity struct {
	*anticheat.Evaluator
	PredictionLatitude float64
	UnstableRate       float64
}

var _ anticheat.MoveDataHandler = &Gravity{}

func init() {
	oldA := Available
	Available = func() []any {
		return append(oldA(), &Gravity{
			anticheat.NewEvaluator(80, 0.75, 0.96),
			0.005,
			0.997,
		})
	}
}

const epsilon = 0.001

func (g *Gravity) HandleMoveData(p *anticheat.InternalPlayer, data *xyron.PlayerMoveData) *xyron.JudgementData {
	tickSinceTeleport := p.CurrentTimestamp() - p.Teleport.Current()
	if p.Motion.Current() != nil {
		return nil
	}
	if p.Location.Previous() == nil ||
		tickSinceTeleport < 15 ||
		//TODO better high jump support
		p.InAirTick < 15 ||
		p.OnGround.Current() {
		return nil
	}
	newOnGround, _, _, _, _ := p.CheckGroundState(data.NewPosition)
	//sometimes when player land and death, false positives appear
	if newOnGround {
		return nil
	}
	if p.Location.Current().IsFlying {
		return nil
	}
	motion := p.Motion.Current()
	if motion != nil {
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
	prevDeltaY := oldPos.Sub(oldOldPos).Y()
	measuredDeltaY := newPos.Sub(oldPos).Y()
	predictedDeltaY := g.predictDeltaY(p, data, prevDeltaY)

	equalness := 1 - math.Min(measuredDeltaY/predictedDeltaY, predictedDeltaY/measuredDeltaY)

	g.HandleMaxRate(equalness, g.PredictionLatitude, g.UnstableRate)
	return &xyron.JudgementData{
		Type:      "Gravity",
		Judgement: g.Evaluate(),
		Message:   fmt.Sprintf("p:%v pred-dy:%.5f dy:%.5f eq:%.5f", g.PossibilityString(), predictedDeltaY, prevDeltaY, equalness),
	}
}

// predictDeltaY https://github.com/Blackjack200/minecraft_client_1_16_2/blob/master/net/minecraft/world/entity/LivingEntity.java#L1891-1911
func (g *Gravity) predictDeltaY(p *anticheat.InternalPlayer, data *xyron.PlayerMoveData, prevDeltaY float64) float64 {
	predictedDeltaY := prevDeltaY
	if e, ok := p.GetEffect(func(f *xyron.EffectFeature) bool {
		return f.IsLevitation
	}); ok {
		predictedDeltaY += (0.05*(float64(e.Amplifier+1)) - prevDeltaY) * 0.2
	} else if p.Location.Current().HaveGravity {
		predictedDeltaY -= calculateGravity(p)
	}
	predictedDeltaY *= 0.9800000190734863

	//Cobweb https://github.com/Blackjack200/minecraft_client_1_16_2/blob/master/net/minecraft/world/entity/Entity.java#L516
	//https://github.com/Blackjack200/minecraft_client_1_16_2/blob/c7f87b96efaeb477d9604354aa23ada0eb637ec6/net/minecraft/world/level/block/WebBlock.java#L17C1
	if p.InCobweb.Current() {
		predictedDeltaY *= 0.05000000074505806
	}

	//SweetBerry https://github.com/Blackjack200/minecraft_client_1_16_2/blob/c7f87b96efaeb477d9604354aa23ada0eb637ec6/net/minecraft/world/level/block/SweetBerryBushBlock.java#L73
	if p.InCobweb.Current() {
		predictedDeltaY *= 0.75
	}
	return predictedDeltaY
}

func calculateGravity(p *anticheat.InternalPlayer) float64 {
	gravity := 0.08
	if !p.OnGround.Current() && p.HasEffect(func(f *xyron.EffectFeature) bool {
		return f.IsSlowFalling
	}) {
		gravity = 0.01
	}
	return gravity
}
