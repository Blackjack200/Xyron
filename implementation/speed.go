package implementation

import (
	"github.com/blackjack200/xyron/anticheat"
	"github.com/blackjack200/xyron/xyron"
	"github.com/go-gl/mathgl/mgl64"
)

type Speed struct {
	*anticheat.Evaluator
	UnstableRate float64
}

var _ anticheat.MoveDataHandler = &Speed{}

func init() {
	Available = append(Available, &Speed{
		anticheat.NewEvaluator(8, 0.3, 0.8),
		0.9999,
	})
}

const epsilon = 0.001

func (s *Speed) HandleMoveData(p *anticheat.InternalPlayer, data *xyron.PlayerMoveData) *xyron.JudgementData {
	if p.GetVolatile().Teleported {
		return nil
	}
	newPos := toVec3(data.NewPosition.Location.Position)
	newRot := toVec3(data.NewPosition.Location.Direction)
	oldPos := toVec3(p.Location.Previous().Location.Position)
	oldRot := toVec3(p.Location.Previous().Location.Direction)
	delta := newPos.Sub(oldPos)
	deltaPlane := mgl64.Vec2{delta.X(), delta.Z()}
	if deltaPlane.LenSqr() < epsilon {
		return nil
	}

}

// Stolen from Artemis Client, could be quite inaccurate.
func calculateSpeedAmplifier(movement float64) {

}
