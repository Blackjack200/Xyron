package anticheat

import (
	"context"
	"fmt"
	"github.com/blackjack200/xyron/xyron"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/emptypb"
	"sync"
	"sync/atomic"
	"time"
)

type SimpleAnticheat struct {
	xyron.AnticheatServer
	mu      *sync.Mutex
	log     *logrus.Logger
	players map[string]*InternalPlayer
	checks  func() []any
	running atomic.Bool
}

func NewSimpleAnticheatServer(log *logrus.Logger, checks func() []any) (*SimpleAnticheat, func()) {
	s := &SimpleAnticheat{
		mu:      &sync.Mutex{},
		log:     log,
		players: make(map[string]*InternalPlayer),
		checks:  checks,
		running: atomic.Bool{},
	}
	s.running.Store(true)
	go func() {
		t := time.NewTicker(time.Second * 5)
		for s.running.Load() {
			select {
			case _ = <-t.C:
				s.mu.Lock()
				for id, p := range s.players {
					if time.Now().Sub(p.lastReport).Seconds() > 30 {
						s.log.
							WithField("player", p.Name).
							WithField("player_id", id).
							Debugf("timeout")
						delete(s.players, id)
					}
				}
				s.mu.Unlock()
			}
			time.Sleep(time.Second)
		}
	}()
	return s, func() {
		s.running.Store(false)
	}
}

func (s *SimpleAnticheat) AddPlayer(_ context.Context, req *xyron.AddPlayerRequest) (*xyron.PlayerReceipt, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.players[req.Player.Name]; ok {
		return nil, fmt.Errorf("player already exists: %v", req.Player.Name)
	}

	id := internalId(req.Player.Name)
	p := NewInternalPlayer(s.log, s.checks(), req.Player.Os, req.Player.Name)
	s.players[id] = p
	s.handleData(p, req.Data)

	s.log.
		WithField("player", req.Player.Name).
		WithField("player_id", id).
		Debugf("add")

	return &xyron.PlayerReceipt{InternalId: req.Player.Name}, nil
}

func (s *SimpleAnticheat) RemovePlayer(_ context.Context, r *xyron.PlayerReceipt) (*emptypb.Empty, error) {
	s.mu.Lock()
	s.log.
		WithField("player_id", r.InternalId).
		Debugf("remove")
	delete(s.players, r.InternalId)
	s.mu.Unlock()
	return &emptypb.Empty{}, nil
}

func (s *SimpleAnticheat) Report(_ context.Context, r *xyron.ReportData) (*xyron.ReportResponse, error) {
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

func (s *SimpleAnticheat) ReportBatched(_ context.Context, data *xyron.BatchedReportData) (*xyron.BatchedReportResponse, error) {
	f := func(d *xyron.ReportData) *xyron.BatchedReportResponseEntry {
		s.mu.Lock()
		var p *InternalPlayer
		if pp, ok := s.players[d.Player.GetInternalId()]; !ok {
			return nil
		} else {
			p = pp
		}
		s.mu.Unlock()
		//log.Printf("RP:%v", r.Player.InternalId)
		jdjm := s.handleData(p, d.Data)
		return &xyron.BatchedReportResponseEntry{
			Player:     d.Player,
			Judgements: jdjm,
		}
	}
	wg := sync.WaitGroup{}
	mu := &sync.Mutex{}
	var res []*xyron.BatchedReportResponseEntry
	for _, d := range data.Data {
		d := d
		wg.Add(1)
		go func() {
			resp := f(d)
			if resp != nil {
				mu.Lock()
				res = append(res, resp)
				mu.Unlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	return &xyron.BatchedReportResponse{Data: nil}, nil
}
