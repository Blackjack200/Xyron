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

func newTickedBool(d bool) *TickedData[bool] {
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
	Location *xyron.EntityPositionData

	Jumped *TickedData[bool]
}

func (i *InternalPlayer) Tick() {
	i.Jumped.Reset()
}

type SimpleAnticheatServer struct {
	xyron.AnticheatServer
	mu      *sync.Mutex
	players map[string]*InternalPlayer
	checks  []AnticheatCheck
}

// AnticheatCheck this is just a simple implementation
type AnticheatCheck func(p *InternalPlayer, timestamp int32, data *xyron.WildcardReportData) *xyron.JudgementData

func (s *SimpleAnticheatServer) handleData(p *InternalPlayer, tdata map[int32]*xyron.TimestampedReportData) (r []*xyron.JudgementData) {
	checks := s.checks
	keys := ComparableSlice[int32](make([]int32, len(tdata)))
	for timestamp, _ := range tdata {
		keys = append(keys, timestamp)
	}
	keys.Sort()
	for _, timestamp := range keys {
		for _, wdata := range tdata[timestamp].Data {
			switch data := wdata.Data.(type) {
			case *xyron.WildcardReportData_ActionData:
				p.Location = data.ActionData.Position
				switch data.ActionData.Action {
				case xyron.PlayerAction_Jump:
					p.Jumped.Set(true)
				}
				//TODO
			case *xyron.WildcardReportData_MoveData:
				p.Location = data.MoveData.NewPosition
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
				//TODO
			case *xyron.WildcardReportData_InputModeData:
				p.Input = data.InputModeData.InputMode
			case *xyron.WildcardReportData_HeldItemChangeData:
			//TODO
			case *xyron.WildcardReportData_ServerTickData:
				//the end of the tick, useless for now but we can make sure everything is OK?
			}
			for _, c := range checks {
				r = append(r, c(p, timestamp, wdata))
			}
		}
		p.Tick()
	}
	return
}

func (s *SimpleAnticheatServer) AddPlayer(_ context.Context, req *xyron.AddPlayerRequest) (*xyron.PlayerReceipt, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.players[req.Player.Name]; ok {
		return &xyron.PlayerReceipt{InternalId: req.Player.Name}, nil
	}
	log.Printf("AP:%v", req.Player.Name)
	ip := &InternalPlayer{
		Os:       req.Player.Os,
		Input:    0,
		Name:     req.Player.Name,
		GameMode: 0,
		Effects:  make(map[string]struct{}),
		Jumped:   newTickedBool(false),
	}
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

var airJump = AnticheatCheck(func(p *InternalPlayer, timestamp int32, wdata *xyron.WildcardReportData) *xyron.JudgementData {
	if wwdata, ok := wdata.GetData().(*xyron.WildcardReportData_ActionData); ok {
		d := wwdata.ActionData
		if d.GetAction() == xyron.PlayerAction_Jump {
			onGround := false
			for _, b := range d.Position.CollidedBlocks {
				if b.Feature.IsSolid {
					onGround = true
					break
				}
			}
			if !onGround {
				return &xyron.JudgementData{
					Type:      "AirJump",
					Judgement: xyron.Judgement_AMBIGUOUS,
					Message:   fmt.Sprintf("onGround=%v", onGround),
				}
			}
		}
	}
	return nil
})

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
		checks: []AnticheatCheck{
			airJump,
		},
	}
	xyron.RegisterAnticheatServer(s, t)
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
	s.Stop()
}
