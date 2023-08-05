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
	register(func() any {
		return &Gravity{
			anticheat.NewEvaluator(80, 0.75, 0.96),
			0.005,
			0.997,
		}
	})
}

func (g *Gravity) HandleMoveData(p *anticheat.InternalPlayer, data *xyron.PlayerMoveData) *xyron.JudgementData {
	if !isPlayerFreeFalling(p, data.NewPosition) {
		return nil
	}
	if p.Location.Previous() == nil {
		return nil
	}
	oldPos := toVec3(p.Location.Previous().Position)
	pos := toVec3(p.Location.Current().Position)
	deltaY := pos.Sub(oldPos).Y()

	futurePos := toVec3(data.NewPosition.Position)
	measuredFutureDeltaY := futurePos.Sub(pos).Y()
	if isZero(futurePos.Sub(pos).Len()) {
		return nil
	}

	predictedDeltaY := g.predictDeltaY(p, deltaY)

	if !p.Location.Current().IsFlying &&
		!data.NewPosition.IsFlying &&
		!p.Location.Current().AllowFlying {
		g.HandleRelativeUnstableRate(measuredFutureDeltaY, predictedDeltaY, g.PredictionLatitude, g.UnstableRate)
	}

	equalness := math.Abs(measuredFutureDeltaY - predictedDeltaY)
	return &xyron.JudgementData{
		Type:      "Gravity",
		Judgement: g.Evaluate(),
		Message:   fmt.Sprintf("p:%v pred-dy:%.5f dy:%.5f delta:%.5f", g.PossibilityString(), predictedDeltaY, deltaY, equalness),
	}
}

// predictDeltaY https://github.com/Blackjack200/minecraft_client_1_16_2/blob/master/net/minecraft/world/entity/LivingEntity.java#L1891-1911
func (g *Gravity) predictDeltaY(p *anticheat.InternalPlayer, prevDeltaY float64) float64 {
	predictedDeltaY := prevDeltaY
	if amp := p.Effect(func(f *xyron.EffectFeature) bool {
		return f.IsLevitation
	}); amp != 0 {
		predictedDeltaY += (0.05*(amp+1) - prevDeltaY) * 0.2
	} else if p.Location.Current().HaveGravity {
		predictedDeltaY -= calculateGravity(p)
	}
	predictedDeltaY *= 0.98

	//FIXME stuck block prediction not works at all

	//Cobweb https://github.com/Blackjack200/minecraft_client_1_16_2/blob/master/net/minecraft/world/entity/Entity.java#L516
	//https://github.com/Blackjack200/minecraft_client_1_16_2/blob/c7f87b96efaeb477d9604354aa23ada0eb637ec6/net/minecraft/world/level/block/WebBlock.java#L17C1
	if p.InCobweb.Current().Get() {
		println("COB")
		predictedDeltaY *= 0.05
	}

	//SweetBerry https://github.com/Blackjack200/minecraft_client_1_16_2/blob/c7f87b96efaeb477d9604354aa23ada0eb637ec6/net/minecraft/world/level/block/SweetBerryBushBlock.java#L73
	if p.InSweetBerry.Current().Get() {
		predictedDeltaY *= 0.75
	}

	return predictedDeltaY
}

func calculateGravity(p *anticheat.InternalPlayer) float64 {
	gravity := 0.08
	if !p.OnGround.Current().Get() && p.Effect(func(f *xyron.EffectFeature) bool {
		return f.IsSlowFalling
	}) != 0 {
		gravity = 0.01
	}
	return gravity
}
