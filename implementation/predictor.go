package implementation

import (
	"github.com/blackjack200/xyron/anticheat"
	"github.com/blackjack200/xyron/xyron"
	"github.com/go-gl/mathgl/mgl64"
	"math"
)

type predictor struct {
	p *anticheat.InternalPlayer
}

const defaultFlyingSpeed = 0.02
const defaultSpeed = 0.10000000149011612

func (p *predictor) calculateFrictionInfluencedSpeed(friction float64) float64 {
	if p.p.OnGround.Previous() {
		if p.p.HasEffect(func(f *xyron.EffectFeature) bool {
			return f.IsSpeed
		}) {
			//TODO defaultSpeed may change
			return defaultSpeed * (0.21600002 / (friction * friction * friction))
		}
	}
	//TODO flyingSpeed may change
	return defaultFlyingSpeed
}

func clamp(num, min, max float64) float64 {
	if num < min {
		return min
	}
	if num > max {
		return max
	}
	return num
}

func (p *predictor) handleOnClimbable(dck mgl64.Vec3) mgl64.Vec3 {
	if p.p.OnClimbable.Previous() {
		dx := clamp(dck.X(), -0.15000000596046448, 0.15000000596046448)
		dz := clamp(dck.Z(), -0.15000000596046448, 0.15000000596046448)

		dy := math.Max(dck.Y(), -0.15000000596046448)
		//TODO !SCAFFOLDING || this.isSuppressingSlidingDownLadder()
		if dy < 0.0 && p.p.Sneaking.Previous() {
			dy = 0.0
		}
		return mgl64.Vec3{dx, dy, dz}
	}
	return dck
}

// getInputVector https://github.com/Blackjack200/minecraft_client_1_16_2/blob/master/net/minecraft/world/entity/Entity.java#L1039
func (p *predictor) getInputVector(dck mgl64.Vec3, frictionSpeed, yRot float64) mgl64.Vec3 {
	double4 := dck.LenSqr()
	if double4 < 1.0e-7 {
		return mgl64.Vec3{0, 0, 0}
	}
	dck2 := dck
	if double4 > 1.0 {
		dck2 = dck.Normalize()
	}
	dck2 = dck2.Mul(frictionSpeed)
	float4 := math.Sin(yRot * 0.017453292)
	float5 := math.Cos(yRot * 0.017453292)
	return mgl64.Vec3{dck2.X()*float5 - dck2.Z()*float4, dck2.Y(), dck2.Z()*float5 + dck2.X()*float4}
}

// handleRelativeFrictionAndCalculateMovement https://github.com/Blackjack200/minecraft_client_1_16_2/blob/master/net/minecraft/world/entity/LivingEntity.java#L1930
// also https://github.com/Blackjack200/minecraft_client_1_16_2/blob/master/net/minecraft/world/entity/Entity.java#L1034
func (p *predictor) handleRelativeFrictionAndCalculateMovement(dck mgl64.Vec3, f float64) mgl64.Vec3 {
	dck2 := p.getInputVector(dck, p.calculateFrictionInfluencedSpeed(f), math.Asin(float64(-p.p.Location.Previous().Location.Direction.Y)))
	dck2 = p.handleOnClimbable(dck2)
	//TODO
	/*if ((this.horizontalCollision || this.jumping) && this.onClimbable()) {
		dck2 = new Vec3(dck2.x, 0.2, dck2.z);
	}*/
	if p.p.OnClimbable.Previous() {
		dck2[1] = 0.2
	}
	return dck2
}

func (p *predictor) calculateGravity() float64 {
	gravity := 0.08
	if !p.p.OnGround.Current() && p.p.HasEffect(func(f *xyron.EffectFeature) bool {
		return f.IsSlowFalling
	}) {
		gravity = 0.01
	}
	return gravity
}

// predictDelta https://github.com/Blackjack200/minecraft_client_1_16_2/blob/master/net/minecraft/world/entity/LivingEntity.java#L1891-1911
func (p *predictor) predictDelta(data *xyron.PlayerMoveData, prevDelta mgl64.Vec3) mgl64.Vec3 {
	friction := float64(p.p.Location.Previous().BelowThatAffectMovement.Feature.Friction)
	horizontalSpeed := 0.91
	if p.p.OnGround.Previous() {
		horizontalSpeed *= friction
	}
	nextTickDelta := p.handleRelativeFrictionAndCalculateMovement(prevDelta, friction)
	nextTickDeltaY := nextTickDelta.Y()
	if e, ok := p.p.GetEffect(func(f *xyron.EffectFeature) bool {
		return f.IsLevitation
	}); ok {
		nextTickDeltaY += (0.05*(float64(e.Amplifier+1)) - nextTickDelta.Y()) * 0.2
		//TODO wrong HaveGravity used
	} else if data.NewPosition.HaveGravity {
		nextTickDeltaY -= p.calculateGravity()
		//TODO WTF https://github.com/Blackjack200/minecraft_client_1_16_2/blob/master/net/minecraft/world/entity/LivingEntity.java#L1905
	} else {
		nextTickDeltaY = 0.0
	}
	nextTickDeltaY *= 0.9800000190734863

	//Cobweb https://github.com/Blackjack200/minecraft_client_1_16_2/blob/master/net/minecraft/world/entity/Entity.java#L516
	//https://github.com/Blackjack200/minecraft_client_1_16_2/blob/c7f87b96efaeb477d9604354aa23ada0eb637ec6/net/minecraft/world/level/block/WebBlock.java#L17C1
	if p.p.InCobweb.Previous() {
		nextTickDeltaY *= 0.05000000074505806
	}

	//SweetBerry https://github.com/Blackjack200/minecraft_client_1_16_2/blob/c7f87b96efaeb477d9604354aa23ada0eb637ec6/net/minecraft/world/level/block/SweetBerryBushBlock.java#L73
	if p.p.InCobweb.Previous() {
		nextTickDeltaY *= 0.75
	}
	return mgl64.Vec3{
		horizontalSpeed * nextTickDelta.X(),
		nextTickDeltaY,
		horizontalSpeed * nextTickDelta.Z(),
	}
}
