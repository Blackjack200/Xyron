package implementation

import (
	"github.com/blackjack200/xyron/anticheat"
	"github.com/blackjack200/xyron/xyron"
	"github.com/go-gl/mathgl/mgl64"
	"github.com/sirupsen/logrus"
	"math"
)

//FIXME the entire of predictor not work properly

type predictor struct {
	log *logrus.Logger
}

// predictDelta https://github.com/Blackjack200/minecraft_client_1_16_2/blob/master/net/minecraft/world/entity/LivingEntity.java#L1891-1911
func (pred *predictor) predictNextTickDeltaMovement(
	p *anticheat.InternalPlayer,
	deltaMovement mgl64.Vec3,
) mgl64.Vec3 {
	newDeltaMovement := deltaMovement
	if p.Swimming.Current().Get() {
		dirc := toVec3(p.Location.Current().Direction)
		dircY := dirc.Y()
		d := 0.06
		if dircY < -0.2 {
			d = 0.085
		}
		if dircY <= 0.0 || p.Jump.Current().Duration(p.CurrentTimestamp()) == 0 || !p.IntersectedLiquid.Current().Get() {
			newDeltaMovement[1] += (dircY - newDeltaMovement.Y()) * d
		}
	}

	if p.Location.Current().IsFlying {
		curtY := newDeltaMovement.Y()
		newDeltaMovement = pred.entityTravel(p, newDeltaMovement)
		newDeltaMovement[1] = curtY * 0.6
	}
	newDeltaMovement = pred.entityTravel(p, newDeltaMovement)
	return newDeltaMovement
}

func (pred *predictor) entityTravel(p *anticheat.InternalPlayer, deltaMovement mgl64.Vec3) mgl64.Vec3 {
	//without collides, deltaPosition is deltaMovement
	gravity := 0.08
	isFalling := deltaMovement.Y() <= 0.0
	if !isFalling && p.HasEffect(func(f *xyron.EffectFeature) bool {
		return f.IsSlowFalling
	}) {
		gravity = 0.01
	}

	if p.IntersectedLiquid.Current().Get() /*&& !canStandOnLiquid*/ {

	}

	friction := float64(p.Location.Current().BelowThatAffectMovement.Feature.Friction)
	horizontalSpeedFriction := 0.91
	if p.OnGround.Current().Get() {
		horizontalSpeedFriction *= friction
	}

	predictedDeltaMovement := pred.handleRelativeFrictionAndCalculateMovement(p, deltaMovement, friction)
	predictedDeltaY := predictedDeltaMovement.Y()

	if e, ok := p.Effect(func(f *xyron.EffectFeature) bool {
		return f.IsLevitation
	}); ok {
		predictedDeltaY += (0.05*(float64(e.Amplifier+1)) - predictedDeltaMovement.Y()) * 0.2
	} else if p.Location.Current().HaveGravity {
		predictedDeltaY -= gravity
		//TODO WTF https://github.com/Blackjack200/minecraft_client_1_16_2/blob/master/net/minecraft/world/entity/LivingEntity.java#L1905
	}

	return mgl64.Vec3{
		horizontalSpeedFriction * predictedDeltaMovement.X(),
		predictedDeltaY * 0.98,
		horizontalSpeedFriction * predictedDeltaMovement.Z(),
	}
}

// handleRelativeFrictionAndCalculateMovement https://github.com/Blackjack200/minecraft_client_1_16_2/blob/master/net/minecraft/world/entity/LivingEntity.java#L1930
// also https://github.com/Blackjack200/minecraft_client_1_16_2/blob/master/net/minecraft/world/entity/Entity.java#L1034
func (pred *predictor) handleRelativeFrictionAndCalculateMovement(
	p *anticheat.InternalPlayer,
	deltaMovement mgl64.Vec3,
	friction float64,
) (newDeltaMovement mgl64.Vec3) {
	newDeltaMovement = pred.moveRelative(
		pred.getFrictionInfluencedSpeed(p, friction),
		deltaMovement,
		toVec3(p.Location.Current().Direction),
	)
	newDeltaMovement = pred.handleOnClimbable(p, newDeltaMovement)
	newDeltaMovement = pred.moveSelf(p, newDeltaMovement)
	//TODO
	/*if ((this.horizontalCollision || this.jumping) && this.onClimbable()) {
		dck2 = new Vec3(dck2.x, 0.2, dck2.z);
	}*/
	//if ((p.horizontalCollision || this.jumping) && (this.onClimbable() || this.getFeetBlockState().is(Blocks.POWDER_SNOW) && PowderSnowBlock.canEntityWalkOnPowderSnow(this))) {
	if p.OnClimbable.Current().Get() {
		newDeltaMovement[1] = 0.2
	}
	return
}

func (pred *predictor) handleOnClimbable(p *anticheat.InternalPlayer, dck mgl64.Vec3) mgl64.Vec3 {
	if p.OnClimbable.Current().Get() {
		dx := mgl64.Clamp(dck.X(), -0.15000000596046448, 0.15000000596046448)
		dz := mgl64.Clamp(dck.Z(), -0.15000000596046448, 0.15000000596046448)

		dy := math.Max(dck.Y(), -0.15000000596046448)
		//TODO !SCAFFOLDING || this.isSuppressingSlidingDownLadder()
		if dy < 0.0 && p.Sneaking.Current().Get() {
			dy = 0.0
		}
		return mgl64.Vec3{dx, dy, dz}
	}
	return dck
}

// moveSelf see Entity.move
func (pred *predictor) moveSelf(p *anticheat.InternalPlayer, deltaMovement mgl64.Vec3) mgl64.Vec3 {
	//Cobweb https://github.com/Blackjack200/minecraft_client_1_16_2/blob/master/net/minecraft/world/entity/Entity.java#L516
	//https://github.com/Blackjack200/minecraft_client_1_16_2/blob/c7f87b96efaeb477d9604354aa23ada0eb637ec6/net/minecraft/world/level/block/WebBlock.java#L17C1
	if p.InCobweb.Current().Get() {
		deltaMovement[0] *= 0.25
		deltaMovement[1] *= 0.05
		deltaMovement[2] *= 0.25
	}

	//SweetBerry https://github.com/Blackjack200/minecraft_client_1_16_2/blob/c7f87b96efaeb477d9604354aa23ada0eb637ec6/net/minecraft/world/level/block/SweetBerryBushBlock.java#L73
	if p.InSweetBerry.Current().Get() {
		deltaMovement[0] *= 0.8
		deltaMovement[1] *= 0.75
		deltaMovement[2] *= 0.8
	}
	return deltaMovement
}

const defaultFlyingSpeed = 0.02
const defaultAbilitiesFlyingSpeed = 0.2
const defaultAttributeSpeed = 0.7

// getFrictionInfluencedSpeed see LivingEntity.getFrictionInfluencedSpeed
func (pred *predictor) getFrictionInfluencedSpeed(p *anticheat.InternalPlayer, friction float64) float64 {
	if p.OnGround.Current().Get() {
		//TODO defaultSpeed may change
		return pred.getSpeed(p) * (0.21600002 / (friction * friction * friction))
	}
	//TODO flyingSpeed may change
	return pred.getFlyingSpeed(p)
}

func (pred *predictor) getSpeed(p *anticheat.InternalPlayer) float64 {
	speed := defaultAttributeSpeed
	if e, ok := p.Effect(func(f *xyron.EffectFeature) bool {
		return f.IsSpeed
	}); ok {
		//modifier
		speed += 0.2 * float64(e.Amplifier)
	}
	return speed
}

// getFlyingSpeed see Player.getFlyingSpeed
func (pred *predictor) getFlyingSpeed(p *anticheat.InternalPlayer) float64 {
	if p.Location.Current().IsFlying {
		if p.Sprinting.Current().Get() {
			return defaultAbilitiesFlyingSpeed * 2
		}
		return defaultAbilitiesFlyingSpeed
	}
	if p.Sprinting.Current().Get() {
		return 0.025999999
	}
	return 0.02
}

func (pred *predictor) moveRelative(friction float64, deltaMovement mgl64.Vec3, direction mgl64.Vec3) mgl64.Vec3 {
	yRot := ToRotation(direction).Pitch()
	inputDelta := pred.getInputVector(deltaMovement, friction, yRot)
	return deltaMovement.Add(inputDelta)
}

// getInputVector https://github.com/Blackjack200/minecraft_client_1_16_2/blob/master/net/minecraft/world/entity/Entity.java#L1039
func (pred *predictor) getInputVector(dck mgl64.Vec3, frictionSpeed, yRot float64) mgl64.Vec3 {
	lenSqr := dck.LenSqr()
	if lenSqr < 1.0e-7 {
		return mgl64.Vec3{0, 0, 0}
	}
	dck2 := dck
	if lenSqr > 1.0 {
		dck2 = dck.Normalize()
	}
	dck2 = dck2.Mul(frictionSpeed)
	float4 := math.Sin(mgl64.DegToRad(yRot))
	float5 := math.Cos(mgl64.DegToRad(yRot))
	return mgl64.Vec3{dck2.X()*float5 - dck2.Z()*float4, dck2.Y(), dck2.Z()*float5 + dck2.X()*float4}
}
