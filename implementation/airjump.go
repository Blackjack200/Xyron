package implementation

import (
	"fmt"
	"github.com/blackjack200/xyron/anticheat"
	"github.com/blackjack200/xyron/xyron"
)

type AirJump struct{}

var _ = anticheat.ActionDataHandler(&AirJump{})

func (a *AirJump) HandleActionData(p *anticheat.InternalPlayer, data *xyron.PlayerActionData) *xyron.JudgementData {
	if data.Action == xyron.PlayerAction_Jump && !p.OnGround.Current() && !p.OnGround.Previous() {
		return &xyron.JudgementData{
			Type:      "AirJump",
			Judgement: xyron.Judgement_AMBIGUOUS,
			Message:   fmt.Sprintf("onGround: cur:%v prev:%v", p.OnGround.Current(), p.OnGround.Previous()),
		}
	}
	return nil
}
