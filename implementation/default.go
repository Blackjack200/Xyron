package implementation

import (
	"github.com/blackjack200/xyron/anticheat"
	"github.com/blackjack200/xyron/xyron"
)

var Available = func() []any {
	return nil
}

const epsilon = 0.00001

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

	// we shouldn't use "future" data, but this is a special condition, false positives appear when player land and death.
	futureOnGround, _, _, _, _, _, _ := p.CheckGroundState(futurePos)
	futureImmobile := futurePos.IsImmobile

	if !isZero(y) &&
		!p.Location.Current().IsImmobile && !futureImmobile &&
		futurePos.Position.Y > -64 &&
		tickSinceTeleport > 40 && tickSinceFlying > 10 && tickSinceJump > 15 &&
		tickSinceMotion > p.MotionCoolDown &&
		p.InAirTick > 15 && !futureOnGround &&
		!p.IntersectedLiquid.Current().Get() {
		return true
	}
	return false
}
