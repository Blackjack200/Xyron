package implementation

import (
	"fmt"
	"github.com/blackjack200/xyron/anticheat"
	"github.com/blackjack200/xyron/xyron"
)

type JumpCondition struct {
	*anticheat.Evaluator
	UnstableRate float64
}

var _ = anticheat.ActionDataHandler(&JumpCondition{})

func init() {
	register(func() any {
		return &JumpCondition{
			anticheat.NewEvaluator(8, 0.3, 0.8),
			0.9999,
		}
	})
}

func (a *JumpCondition) HandleActionData(p *anticheat.InternalPlayer, data *xyron.PlayerActionData) *xyron.JudgementData {
	measured := 0.0
	newOnGround, _, _, _, _, _ := p.CheckGroundState(data.Position)
	if data.Action == xyron.PlayerAction_Jump &&
		!p.OnGround.Current().Get() &&
		!newOnGround &&
		p.InAirTick >= 15 {
		measured = 1
	}
	if !p.Location.Current().AllowFlying && !data.Position.AllowFlying {
		a.HandleUnstableRate(measured, 0, a.UnstableRate)
	}
	return &xyron.JudgementData{
		Type:      "JumpCondition",
		Judgement: a.Evaluate(),
		Message:   fmt.Sprintf("p:%v inAirTick:%v", a.PossibilityString(), p.InAirTick),
	}
}
