package implementation

import (
	"fmt"
	"github.com/blackjack200/xyron/anticheat"
	"github.com/blackjack200/xyron/xyron"
)

type AirJump struct {
	*anticheat.Evaluator
	UnstableRate float64
}

var _ = anticheat.ActionDataHandler(&AirJump{})

func init() {
	oldA := Available
	Available = func() []any {
		return append(oldA(), &AirJump{
			anticheat.NewEvaluator(8, 0.3, 0.8),
			0.9999,
		})
	}
}

func (a *AirJump) HandleActionData(p *anticheat.InternalPlayer, data *xyron.PlayerActionData) *xyron.JudgementData {
	measured := 0.0
	newOnGround, _, _, _, _ := p.CheckGroundState(data.Position)
	if data.Action == xyron.PlayerAction_Jump &&
		!p.OnGround.Current() &&
		!newOnGround &&
		p.InAirTick >= 15 {
		measured = 1
	}
	a.HandleUnstableRate(measured, 0, a.UnstableRate)
	return &xyron.JudgementData{
		Type:      "AirJump",
		Judgement: a.Evaluate(),
		Message:   fmt.Sprintf("p:%v inAirTick:%v", a.PossibilityString(), p.InAirTick),
	}
}
