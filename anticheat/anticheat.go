package anticheat

import (
	"context"
	"fmt"
	"github.com/blackjack200/xyron/xyron"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"sync"
)

type SimpleAnticheat struct {
	xyron.AnticheatServer
	mu      *sync.Mutex
	log     *logrus.Logger
	players map[string]*InternalPlayer
	checks  func() []any
}

func NewSimpleAnticheatServer(log *logrus.Logger, checks func() []any) *SimpleAnticheat {
	return &SimpleAnticheat{
		mu:      &sync.Mutex{},
		log:     log,
		players: make(map[string]*InternalPlayer),
		checks:  checks,
	}
}

func (s *SimpleAnticheat) AddPlayer(_ context.Context, req *xyron.AddPlayerRequest) (*xyron.PlayerReceipt, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.players[req.Player.Name]; ok {
		return &xyron.PlayerReceipt{InternalId: req.Player.Name}, nil
	}
	log.Printf("AP:%v", req.Player.Name)
	ip := NewInternalPlayer(s.log, s.checks(), req.Player.Os, req.Player.Name)
	s.players[req.Player.Name] = ip
	s.handleData(ip, req.Data)
	return &xyron.PlayerReceipt{InternalId: req.Player.Name}, nil
}

func (s *SimpleAnticheat) RemovePlayer(_ context.Context, r *xyron.PlayerReceipt) (*emptypb.Empty, error) {
	s.mu.Lock()
	log.Printf("DP:%v", r.InternalId)
	delete(s.players, r.InternalId)
	s.mu.Unlock()
	return &emptypb.Empty{}, nil
}

func (s *SimpleAnticheat) Report(_ context.Context, r *xyron.PlayerReport) (*xyron.ReportResponse, error) {
	var p *InternalPlayer
	s.mu.Lock()
	if pp, ok := s.players[r.Player.GetInternalId()]; !ok {
		return nil, fmt.Errorf("player %v not found", r.Player.InternalId)
	} else {
		p = pp
	}
	s.mu.Unlock()
	//log.Printf("RP:%v", r.Player.InternalId)
	jdjm := s.handleData(p, r.Data)
	return &xyron.ReportResponse{Judgements: jdjm}, nil
}
