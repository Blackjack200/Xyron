package anticheat

import (
	"github.com/blackjack200/xyron/xyron"
	"time"
)

func (s *SimpleAnticheat) handleData(p *InternalPlayer, tdata map[int64]*xyron.TimestampedReportData) (r []*xyron.JudgementData) {
	var keys []int64
	for timestamp, _ := range tdata {
		keys = append(keys, timestamp)
	}
	sorted := ComparableSlice[int64](keys)
	sorted.Sort()
	p.lastReport = time.Now()
	for _, timestamp := range sorted {
		p.timestampThisTick = timestamp
		for _, wdata := range tdata[timestamp].Data {
			for _, c := range p.checks {
				c := c
				data := s.callHandlers(p, c, wdata)
				r = append(r, data...)
			}
			s.tickPlayer(p, wdata)
		}
		p.Tick()
	}
	return
}

func (s *SimpleAnticheat) tickPlayer(p *InternalPlayer, wdata *xyron.WildcardReportData) {
	switch data := wdata.Data.(type) {
	case *xyron.WildcardReportData_ActionData:
		p.SetLocation(data.ActionData.Position)
		switch data.ActionData.Action {
		case xyron.PlayerAction_Jump:
			p.Jump.Set(p.timestampThisTick, p.Effect(func(f *xyron.EffectFeature) bool {
				return f.IsJumpBoost
			}))
		case xyron.PlayerAction_StartSprint:
			p.Sprinting.Set(p.timestampThisTick, true)
		case xyron.PlayerAction_StopSprint:
			p.Sprinting.Set(p.timestampThisTick, false)
		case xyron.PlayerAction_StartSneak:
			p.Sneaking.Set(p.timestampThisTick, true)
		case xyron.PlayerAction_StopSneak:
			p.Sneaking.Set(p.timestampThisTick, false)
		case xyron.PlayerAction_StartGliding:
			p.Gliding.Set(p.timestampThisTick, true)
		case xyron.PlayerAction_StopGliding:
			p.Gliding.Set(p.timestampThisTick, false)
		case xyron.PlayerAction_StartSwimming:
			p.Swimming.Set(p.timestampThisTick, true)
		case xyron.PlayerAction_StopSwimming:
			p.Swimming.Set(p.timestampThisTick, false)
		case xyron.PlayerAction_StartSprintFlying:
			p.Flying.Set(p.timestampThisTick, true)
		case xyron.PlayerAction_StopSprintFlying:
			p.Flying.Set(p.timestampThisTick, false)
		case xyron.PlayerAction_OpenInventory:
			p.OpenInventory.Set(p.timestampThisTick, true)
		case xyron.PlayerAction_CloseInventory:
			p.CloseInventory.Set(p.timestampThisTick, true)
		default:
			s.log.Errorf("unhandled action data: %v", data.ActionData.Action)
		}
	case *xyron.WildcardReportData_MoveData:
		p.SetLocation(data.MoveData.NewPosition)
		if data.MoveData.Teleport {
			p.Teleport.Set(p.timestampThisTick, toVec3(data.MoveData.NewPosition.Position))
		}
	case *xyron.WildcardReportData_PlaceBlockData:
		p.PlaceBlock.Set(p.timestampThisTick, data.PlaceBlockData)
	case *xyron.WildcardReportData_BreakBlockData:
		p.BreakBlock.Set(p.timestampThisTick, data.BreakBlockData)
	case *xyron.WildcardReportData_EatFoodData:
		p.Eat.Set(p.timestampThisTick, data.EatFoodData.Status == xyron.ConsumeStatus_Stop)
	case *xyron.WildcardReportData_AttackData:
		p.Attack.Set(p.timestampThisTick, data.AttackData.Data)
	case *xyron.WildcardReportData_EffectData:
		p.effects = data.EffectData.Effect
	case *xyron.WildcardReportData_GameModeData:
		p.GameMode = data.GameModeData.GameMode
	case *xyron.WildcardReportData_MotionData:
		motion := data.MotionData.Motion
		p.Motion.Set(p.timestampThisTick, toVec3(motion))
		p.MotionCoolDown += int64(((motion.X+motion.Y)/2 + 2) * 15)
	case *xyron.WildcardReportData_InputModeData:
		p.Input = data.InputModeData.InputMode
	case *xyron.WildcardReportData_LifeData:
		p.Alive.Set(p.timestampThisTick, data.LifeData.Alive)
	case *xyron.WildcardReportData_HeldItemChangeData:
		p.HeldItem.Set(p.timestampThisTick, data.HeldItemChangeData.Item)
	default:
		s.log.Errorf("unhandled data: %T", data)
	}
}

// ActionDataHandler handles *xyron.WildcardReportData_ActionData
type ActionDataHandler interface {
	HandleActionData(*InternalPlayer, *xyron.PlayerActionData) *xyron.JudgementData
}

// MoveDataHandler handles *xyron.WildcardReportData_MoveData
type MoveDataHandler interface {
	HandleMoveData(*InternalPlayer, *xyron.PlayerMoveData) *xyron.JudgementData
}

// PlaceBlockDataHandler handles *xyron.WildcardReportData_PlaceBlockData
type PlaceBlockDataHandler interface {
	HandlePlaceBlockData(*InternalPlayer, *xyron.PlayerPlaceBlockData) *xyron.JudgementData
}

// BreakBlockDataHandler handles *xyron.WildcardReportData_BreakBlockData
type BreakBlockDataHandler interface {
	HandleBreakBlockData(*InternalPlayer, *xyron.PlayerBreakBlockData) *xyron.JudgementData
}

// EatFoodDataHandler handles *xyron.WildcardReportData_EatFoodData
type EatFoodDataHandler interface {
	HandleEatFoodData(*InternalPlayer, *xyron.PlayerEatFoodData) *xyron.JudgementData
}

// AttackDataHandler handles *xyron.WildcardReportData_AttackData
type AttackDataHandler interface {
	HandleAttackData(*InternalPlayer, *xyron.PlayerAttackData) *xyron.JudgementData
}

// EffectDataHandler handles *xyron.WildcardReportData_EffectData
type EffectDataHandler interface {
	HandleEffectData(*InternalPlayer, *xyron.PlayerEffectData) *xyron.JudgementData
}

// GameModeDataHandler handles *xyron.WildcardReportData_GameModeData
type GameModeDataHandler interface {
	HandleGameModeData(*InternalPlayer, *xyron.PlayerGameModeData) *xyron.JudgementData
}

// MotionDataHandler handles *xyron.WildcardReportData_MotionData
type MotionDataHandler interface {
	HandleMotionData(*InternalPlayer, *xyron.PlayerMotionData) *xyron.JudgementData
}

// InputModeDataHandler handles *xyron.WildcardReportData_InputModeData
type InputModeDataHandler interface {
	HandleInputModeData(*InternalPlayer, *xyron.PlayerInputModeData) *xyron.JudgementData
}

// HeldItemChangeDataHandler handles *xyron.WildcardReportData_HeldItemChangeData
type HeldItemChangeDataHandler interface {
	HandleHeldItemChangeData(*InternalPlayer, *xyron.PlayerHeldItemChangeData) *xyron.JudgementData
}

func (s *SimpleAnticheat) callHandlers(p *InternalPlayer, c any, wdata *xyron.WildcardReportData) []*xyron.JudgementData {
	var r []*xyron.JudgementData
	if handler, ok := c.(ActionDataHandler); ok {
		if data, ok := wdata.Data.(*xyron.WildcardReportData_ActionData); ok {
			r = append(r, handler.HandleActionData(p, data.ActionData))
		}
	}
	if handler, ok := c.(MoveDataHandler); ok {
		if data, ok := wdata.Data.(*xyron.WildcardReportData_MoveData); ok {
			r = append(r, handler.HandleMoveData(p, data.MoveData))
		}
	}
	if handler, ok := c.(PlaceBlockDataHandler); ok {
		if data, ok := wdata.Data.(*xyron.WildcardReportData_PlaceBlockData); ok {
			r = append(r, handler.HandlePlaceBlockData(p, data.PlaceBlockData))
		}
	}
	if handler, ok := c.(BreakBlockDataHandler); ok {
		if data, ok := wdata.Data.(*xyron.WildcardReportData_BreakBlockData); ok {
			r = append(r, handler.HandleBreakBlockData(p, data.BreakBlockData))
		}
	}
	if handler, ok := c.(EatFoodDataHandler); ok {
		if data, ok := wdata.Data.(*xyron.WildcardReportData_EatFoodData); ok {
			r = append(r, handler.HandleEatFoodData(p, data.EatFoodData))
		}
	}
	if handler, ok := c.(AttackDataHandler); ok {
		if data, ok := wdata.Data.(*xyron.WildcardReportData_AttackData); ok {
			r = append(r, handler.HandleAttackData(p, data.AttackData))
		}
	}
	if handler, ok := c.(EffectDataHandler); ok {
		if data, ok := wdata.Data.(*xyron.WildcardReportData_EffectData); ok {
			r = append(r, handler.HandleEffectData(p, data.EffectData))
		}
	}
	if handler, ok := c.(GameModeDataHandler); ok {
		if data, ok := wdata.Data.(*xyron.WildcardReportData_GameModeData); ok {
			r = append(r, handler.HandleGameModeData(p, data.GameModeData))
		}
	}
	if handler, ok := c.(MotionDataHandler); ok {
		if data, ok := wdata.Data.(*xyron.WildcardReportData_MotionData); ok {
			r = append(r, handler.HandleMotionData(p, data.MotionData))
		}
	}
	if handler, ok := c.(InputModeDataHandler); ok {
		if data, ok := wdata.Data.(*xyron.WildcardReportData_InputModeData); ok {
			r = append(r, handler.HandleInputModeData(p, data.InputModeData))
		}
	}
	if handler, ok := c.(HeldItemChangeDataHandler); ok {
		if data, ok := wdata.Data.(*xyron.WildcardReportData_HeldItemChangeData); ok {
			r = append(r, handler.HandleHeldItemChangeData(p, data.HeldItemChangeData))
		}
	}

	var result []*xyron.JudgementData

	for _, item := range r {
		if item != nil {
			result = append(result, item)
		}
	}
	return result
}
