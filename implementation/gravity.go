package implementation

import (
	"fmt"
	"github.com/blackjack200/xyron/anticheat"
	"github.com/blackjack200/xyron/xyron"
	"math"
)

type Gravity struct {
	*anticheat.Evaluator
	MaxEqualness float64
	UnstableRate float64
}

var _ anticheat.MoveDataHandler = &Gravity{}

func init() {
	Available = append(Available, &Gravity{
		anticheat.NewEvaluator(20, 0.6, 0.9),
		0.05,
		0.9999999999999999,
	})
}

const epsilon = 0.001

func (g *Gravity) HandleMoveData(p *anticheat.InternalPlayer, data *xyron.PlayerMoveData) *xyron.JudgementData {
	tickSinceTeleport := p.CurrentTimestamp() - p.Teleport.Current()
	if p.Motion.Current() == nil {
		return nil
	}
	if p.GetVolatile().Teleported ||
		p.DeltaPosition.Previous() == nil ||
		p.DeltaPosition.Current() == nil ||
		p.Location.Previous() == nil ||
		tickSinceTeleport < 15 ||
		//TODO better high jump support
		p.InAirTick < 15 ||
		p.Location.Previous().Location.Position.Y < -10 ||
		p.OnGround.Current() ||
		p.OnGround.Previous() {
		return nil
	}
	motion := p.Motion.Current()
	maxTick := int64(10)
	//TODO improve big motion support
	if toVec3(motion.Get()).Len() >= 1.5 {
		maxTick = 35
	}
	tickSinceMotion := p.CurrentTimestamp() - motion.Timestamp()
	if tickSinceMotion < maxTick {
		return nil
	}

	prevDeltaY := toVec3(p.DeltaPosition.Previous()).Y()
	deltaY := toVec3(p.DeltaPosition.Current()).Y()

	predictedDeltaY := g.predictDeltaY(p, data, prevDeltaY)

	equalness := math.Abs(predictedDeltaY - deltaY)

	if equalness > g.MaxEqualness {
		g.HandleMaxRate(equalness, g.MaxEqualness, g.UnstableRate)
		return &xyron.JudgementData{
			Type:      "Gravity",
			Judgement: g.Evaluate(),
			Message:   fmt.Sprintf("p:%v pred-dy:%.5f dy:%.5f eq:%.5f", g.PossibilityString(), predictedDeltaY, deltaY, equalness),
		}
	}

	return nil
}

// predictDeltaY https://github.com/Blackjack200/minecraft_client_1_16_2/blob/master/net/minecraft/world/entity/LivingEntity.java#L1891-1911
func (g *Gravity) predictDeltaY(p *anticheat.InternalPlayer, data *xyron.PlayerMoveData, prevDeltaY float64) float64 {
	predictedDeltaY := prevDeltaY
	if e, ok := p.GetEffect(func(f *xyron.EffectFeature) bool {
		return f.IsLevitation
	}); ok {
		predictedDeltaY += (0.05*(float64(e.Amplifier+1)) - prevDeltaY) * 0.2
		//TODO wrong HaveGravity used
	} else if data.NewPosition.HaveGravity {
		predictedDeltaY -= calculateGravity(p)
		//TODO WTF https://github.com/Blackjack200/minecraft_client_1_16_2/blob/master/net/minecraft/world/entity/LivingEntity.java#L1905
	} else if p.Location.Previous().Location.Position.Y > 0.0 {
		predictedDeltaY = -0.1
	} else {
		predictedDeltaY = 0.0
	}
	predictedDeltaY += 0.9800000190734863

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
