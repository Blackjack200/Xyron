package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/blackjack200/xyron/xyron"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"net"
	"sort"
	"sync"
)

type BufferedData[T any] struct {
	prev, cur T
}

func NewBufferedData[T any](cur T) *BufferedData[T] {
	return &BufferedData[T]{prev: cur, cur: cur}
}

func (b *BufferedData[T]) Previous() T {
	return b.prev
}

func (b *BufferedData[T]) Current() T {
	return b.Current()
}

func (b *BufferedData[T]) Set(v T) {
	b.prev, b.cur = b.cur, v
}

type TickedData[T any] struct {
	dv T
	v  T
}

func (t *TickedData[T]) Default(dv T) {
	t.dv = dv
}

func (t *TickedData[T]) Set(v T) {
	t.v = v
}
func (t *TickedData[T]) Get() T {
	return t.v
}

func (t *TickedData[T]) Reset() {
	t.v = t.dv
}

func NewTickedData[T any](dv T) *TickedData[T] {
	return &TickedData[T]{dv: dv, v: dv}
}

func NewTickedBool(d bool) *TickedData[bool] {
	return &TickedData[bool]{
		dv: d,
		v:  d,
	}
}

type ComparableSlice[T ~int |
	~int8 | ~int32 | ~int64 |
	~uint8 | ~uint32 | ~uint64 |
	~float32 | ~float64,
] []T

func (x ComparableSlice[T]) Len() int           { return len(x) }
func (x ComparableSlice[T]) Less(i, j int) bool { return x[i] < x[j] }
func (x ComparableSlice[T]) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

// Sort is a convenience method: x.Sort() calls Sort(x).
func (x ComparableSlice[T]) Sort() { sort.Sort(x) }

type InternalPlayer struct {
	Os       xyron.DeviceOS
	Input    xyron.InputMode
	Name     string
	GameMode xyron.GameMode
	Effects  map[string]struct{}
	Motion   *BufferedData[*xyron.Vec3F]
	location *BufferedData[*xyron.EntityPositionData]
	Volatile *TickedData[*VolatileData]

	OnGround  bool
	InAirTick uint32
}

func NewInternalPlayer(os xyron.DeviceOS, name string) *InternalPlayer {
	return &InternalPlayer{
		Os:       os,
		Name:     name,
		GameMode: 0,
		Effects:  make(map[string]struct{}),
		Motion:   NewBufferedData[*xyron.Vec3F](nil),
		location: NewBufferedData[*xyron.EntityPositionData](nil),
		Volatile: NewTickedData(&VolatileData{}),
		OnGround: true,
	}
}

func (p *InternalPlayer) SetLocation(pos *xyron.EntityPositionData) {
	p.location.Set(pos)
	if pos != nil {
		check := func(bb []*xyron.BlockData) bool {
			for _, b := range bb {
				if b.Feature.IsSolid {
					return true
				}
			}
			return false
		}
		p.OnGround = check(pos.CollidedBlocks) || check(pos.IntersectedBlocks)
	}
}

type VolatileData struct {
	Jumped bool
}

func (p *InternalPlayer) Tick() {
	p.Volatile.Reset()
	p.Motion.Set(nil)
	if !p.OnGround {
		p.InAirTick++
	} else {
		p.InAirTick = 0
	}
}

type SimpleAnticheatServer struct {
	xyron.AnticheatServer
	mu      *sync.Mutex
	players map[string]*InternalPlayer
	checks  []any
}

func NewSimpleAnticheatServer(checks []any) *SimpleAnticheatServer {
	return &SimpleAnticheatServer{
		mu:      &sync.Mutex{},
		players: make(map[string]*InternalPlayer),
		checks:  checks,
	}
}

// ActionDataHandler handles *xyron.WildcardReportData_ActionData
type ActionDataHandler interface {
	handleActionData(*InternalPlayer, *xyron.PlayerActionData) *xyron.JudgementData
}

// MoveDataHandler handles *xyron.WildcardReportData_MoveData
type MoveDataHandler interface {
	handleMoveData(*InternalPlayer, *xyron.PlayerMoveData) *xyron.JudgementData
}

// PlaceBlockDataHandler handles *xyron.WildcardReportData_PlaceBlockData
type PlaceBlockDataHandler interface {
	handlePlaceBlockData(*InternalPlayer, *xyron.PlayerPlaceBlockData) *xyron.JudgementData
}

// BreakBlockDataHandler handles *xyron.WildcardReportData_BreakBlockData
type BreakBlockDataHandler interface {
	handleBreakBlockData(*InternalPlayer, *xyron.PlayerBreakBlockData) *xyron.JudgementData
}

// EatFoodDataHandler handles *xyron.WildcardReportData_EatFoodData
type EatFoodDataHandler interface {
	handleEatFoodData(*InternalPlayer, *xyron.PlayerEatFoodData) *xyron.JudgementData
}

// AttackDataHandler handles *xyron.WildcardReportData_AttackData
type AttackDataHandler interface {
	handleAttackData(*InternalPlayer, *xyron.PlayerAttackData) *xyron.JudgementData
}

// AddEffectDataHandler handles *xyron.WildcardReportData_AddEffectData
type AddEffectDataHandler interface {
	handleAddEffectData(*InternalPlayer, *xyron.PlayerAddEffectData) *xyron.JudgementData
}

// RemoveEffectDataHandler handles *xyron.WildcardReportData_RemoveEffectData
type RemoveEffectDataHandler interface {
	handleRemoveEffectData(*InternalPlayer, *xyron.PlayerRemoveEffectData) *xyron.JudgementData
}

// GameModeDataHandler handles *xyron.WildcardReportData_GameModeData
type GameModeDataHandler interface {
	handleGameModeData(*InternalPlayer, *xyron.PlayerGameModeData) *xyron.JudgementData
}

// MotionDataHandler handles *xyron.WildcardReportData_MotionData
type MotionDataHandler interface {
	handleMotionData(*InternalPlayer, *xyron.PlayerMotionData) *xyron.JudgementData
}

// InputModeDataHandler handles *xyron.WildcardReportData_InputModeData
type InputModeDataHandler interface {
	handleInputModeData(*InternalPlayer, *xyron.PlayerInputModeData) *xyron.JudgementData
}

// HeldItemChangeDataHandler handles *xyron.WildcardReportData_HeldItemChangeData
type HeldItemChangeDataHandler interface {
	handleHeldItemChangeData(*InternalPlayer, *xyron.PlayerHeldItemChangeData) *xyron.JudgementData
}

// ServerTickDataHandler handles *xyron.WildcardReportData_ServerTickData
type ServerTickDataHandler interface {
	handleServerTickData(*InternalPlayer, *xyron.ServerTickData) *xyron.JudgementData
}

func (s *SimpleAnticheatServer) callHandlers(p *InternalPlayer, c any, wdata *xyron.WildcardReportData) []*xyron.JudgementData {
	var r []*xyron.JudgementData
	if handler, ok := c.(ActionDataHandler); ok {
		if data, ok := wdata.Data.(*xyron.WildcardReportData_ActionData); ok {
			r = append(r, handler.handleActionData(p, data.ActionData))
		}
	}
	if handler, ok := c.(MoveDataHandler); ok {
		if data, ok := wdata.Data.(*xyron.WildcardReportData_MoveData); ok {
			r = append(r, handler.handleMoveData(p, data.MoveData))
		}
	}
	if handler, ok := c.(PlaceBlockDataHandler); ok {
		if data, ok := wdata.Data.(*xyron.WildcardReportData_PlaceBlockData); ok {
			r = append(r, handler.handlePlaceBlockData(p, data.PlaceBlockData))
		}
	}
	if handler, ok := c.(BreakBlockDataHandler); ok {
		if data, ok := wdata.Data.(*xyron.WildcardReportData_BreakBlockData); ok {
			r = append(r, handler.handleBreakBlockData(p, data.BreakBlockData))
		}
	}
	if handler, ok := c.(EatFoodDataHandler); ok {
		if data, ok := wdata.Data.(*xyron.WildcardReportData_EatFoodData); ok {
			r = append(r, handler.handleEatFoodData(p, data.EatFoodData))
		}
	}
	if handler, ok := c.(AttackDataHandler); ok {
		if data, ok := wdata.Data.(*xyron.WildcardReportData_AttackData); ok {
			r = append(r, handler.handleAttackData(p, data.AttackData))
		}
	}
	if handler, ok := c.(AddEffectDataHandler); ok {
		if data, ok := wdata.Data.(*xyron.WildcardReportData_AddEffectData); ok {
			r = append(r, handler.handleAddEffectData(p, data.AddEffectData))
		}
	}
	if handler, ok := c.(RemoveEffectDataHandler); ok {
		if data, ok := wdata.Data.(*xyron.WildcardReportData_RemoveEffectData); ok {
			r = append(r, handler.handleRemoveEffectData(p, data.RemoveEffectData))
		}
	}
	if handler, ok := c.(GameModeDataHandler); ok {
		if data, ok := wdata.Data.(*xyron.WildcardReportData_GameModeData); ok {
			r = append(r, handler.handleGameModeData(p, data.GameModeData))
		}
	}
	if handler, ok := c.(MotionDataHandler); ok {
		if data, ok := wdata.Data.(*xyron.WildcardReportData_MotionData); ok {
			r = append(r, handler.handleMotionData(p, data.MotionData))
		}
	}
	if handler, ok := c.(InputModeDataHandler); ok {
		if data, ok := wdata.Data.(*xyron.WildcardReportData_InputModeData); ok {
			r = append(r, handler.handleInputModeData(p, data.InputModeData))
		}
	}
	if handler, ok := c.(HeldItemChangeDataHandler); ok {
		if data, ok := wdata.Data.(*xyron.WildcardReportData_HeldItemChangeData); ok {
			r = append(r, handler.handleHeldItemChangeData(p, data.HeldItemChangeData))
		}
	}
	if handler, ok := c.(ServerTickDataHandler); ok {
		if data, ok := wdata.Data.(*xyron.WildcardReportData_ServerTickData); ok {
			r = append(r, handler.handleServerTickData(p, data.ServerTickData))
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

func (s *SimpleAnticheatServer) handleData(p *InternalPlayer, tdata map[int32]*xyron.TimestampedReportData) (r []*xyron.JudgementData) {
	checks := s.checks
	keys := ComparableSlice[int32](make([]int32, len(tdata)))
	for timestamp, _ := range tdata {
		keys = append(keys, timestamp)
	}
	keys.Sort()
	for _, timestamp := range keys {
		if _, ok := tdata[timestamp]; !ok {
			continue
		}
		for _, wdata := range tdata[timestamp].Data {
			switch data := wdata.Data.(type) {
			case *xyron.WildcardReportData_ActionData:
				p.SetLocation(data.ActionData.Position)
				switch data.ActionData.Action {
				case xyron.PlayerAction_Jump:
					p.Volatile.Get().Jumped = true
				}
				//TODO
			case *xyron.WildcardReportData_MoveData:
				p.SetLocation(data.MoveData.NewPosition)
			case *xyron.WildcardReportData_PlaceBlockData:
				//TODO
			case *xyron.WildcardReportData_BreakBlockData:
				//TODO
			case *xyron.WildcardReportData_EatFoodData:
				//TODO
			case *xyron.WildcardReportData_AttackData:
				//TODO
			case *xyron.WildcardReportData_AddEffectData:
				p.Effects[data.AddEffectData.Effect] = struct{}{}
			case *xyron.WildcardReportData_RemoveEffectData:
				delete(p.Effects, data.RemoveEffectData.Effect)
			case *xyron.WildcardReportData_GameModeData:
				p.GameMode = data.GameModeData.GameMode
			case *xyron.WildcardReportData_MotionData:
				p.Motion.Set(data.MotionData.Motion)
				//TODO
			case *xyron.WildcardReportData_InputModeData:
				p.Input = data.InputModeData.InputMode
			case *xyron.WildcardReportData_HeldItemChangeData:
			//TODO
			case *xyron.WildcardReportData_ServerTickData:
				//the end of the tick, useless for now but we can make sure everything is OK?
			}
			for _, c := range checks {
				r = append(r, s.callHandlers(p, c, wdata)...)
			}
		}
		p.Tick()
	}
	return
}

type Check interface {
}

func (s *SimpleAnticheatServer) AddPlayer(_ context.Context, req *xyron.AddPlayerRequest) (*xyron.PlayerReceipt, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.players[req.Player.Name]; ok {
		return &xyron.PlayerReceipt{InternalId: req.Player.Name}, nil
	}
	log.Printf("AP:%v", req.Player.Name)
	ip := NewInternalPlayer(req.Player.Os, req.Player.Name)
	s.players[req.Player.Name] = ip
	s.handleData(ip, req.Data)
	return &xyron.PlayerReceipt{InternalId: req.Player.Name}, nil
}

func (s *SimpleAnticheatServer) RemovePlayer(_ context.Context, r *xyron.PlayerReceipt) (*emptypb.Empty, error) {
	s.mu.Lock()
	log.Printf("DP:%v", r.InternalId)
	delete(s.players, r.InternalId)
	s.mu.Unlock()
	return &emptypb.Empty{}, nil
}

func (s *SimpleAnticheatServer) Report(_ context.Context, r *xyron.PlayerReport) (*xyron.ReportResponse, error) {
	var p *InternalPlayer
	s.mu.Lock()
	if pp, ok := s.players[r.Player.GetInternalId()]; !ok {
		return nil, fmt.Errorf("player %v not found", r.Player.InternalId)
	} else {
		p = pp
	}
	s.mu.Unlock()
	jdjm := s.handleData(p, r.Data)
	return &xyron.ReportResponse{Judgements: jdjm}, nil
}

type AirJump struct {
}

func (a *AirJump) handleActionData(p *InternalPlayer, data *xyron.PlayerActionData) *xyron.JudgementData {
	if data.Action == xyron.PlayerAction_Jump && !p.OnGround {
		return &xyron.JudgementData{
			Type:      "AirJump",
			Judgement: xyron.Judgement_AMBIGUOUS,
			Message:   fmt.Sprintf("onGround=%v", p.OnGround),
		}
	}
	return nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8888))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Print("Listen")
	s := grpc.NewServer()
	t := &SimpleAnticheatServer{
		mu:      &sync.Mutex{},
		players: make(map[string]*InternalPlayer),
		checks: []any{
			&AirJump{},
		},
	}
	xyron.RegisterAnticheatServer(s, t)
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
	s.Stop()
}
