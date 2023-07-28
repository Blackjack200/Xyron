package implementation

import (
	"github.com/blackjack200/xyron/anticheat"
	"github.com/blackjack200/xyron/xyron"
)

var Available = func() []any {
	return nil
}

const epsilon = 0.00000001

func register(newF func() any) {
	oldA := Available
	Available = func() []any {
		return append(oldA(), newF())
	}
}

func clear() {
	Available = func() []any {
		return nil
	}
}

func isPlayerFreeFalling(p *anticheat.InternalPlayer, futurePos *xyron.EntityPositionData) bool {
	if p.Location.Previous() == nil {
		return true
	}
	y := p.Motion.Current().Get().Y()

	tickSinceTeleport := p.Teleport.Current().Duration(p.CurrentTimestamp())
	tickSinceFlying := p.Flying.Current().Duration(p.CurrentTimestamp())
	tickSinceMotion := p.Motion.Current().Duration(p.CurrentTimestamp())
	tickSinceJump := p.Jump.Current().Duration(p.CurrentTimestamp())
	motionDelay := int64(0)
	// we shouldn't use "future" data, but this is a special condition, false positives appear when player land and death.
	futureOnGround, _, _, _, _, _ := p.CheckGroundState(futurePos)
	futureImmobile := futurePos.IsImmobile

	motion := p.Motion.Current()
	if !isZero(motion.Get().Len()) {
		motionDelay = int64(10)
		//TODO improve big motion support
		if motion.Get().Len() >= 1.5 {
			motionDelay = 35
		}
	}

	if !isZero(y) &&
		!p.Location.Current().IsImmobile && !futureImmobile &&
		tickSinceTeleport > 40 && tickSinceFlying > 10 &&
		tickSinceMotion > motionDelay && tickSinceJump > 15 &&
		p.InAirTick > 15 && !futureOnGround &&
		!p.IntersectedLiquid.Current().Get() {
		return true
	}
	return false
}
