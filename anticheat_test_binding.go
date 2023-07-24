package main

import (
	"context"
	"fmt"
	"github.com/blackjack200/xyron/anticheat"
	"github.com/blackjack200/xyron/implementation"
	"github.com/blackjack200/xyron/xyron"
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/block"
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/block/model"
	"github.com/df-mc/dragonfly/server/entity/effect"
	"github.com/df-mc/dragonfly/server/event"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/chat"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/go-gl/mathgl/mgl64"
	"github.com/pelletier/go-toml"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"math"
	"os"
	"reflect"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

type WrappedAnticheatClient struct {
	server *anticheat.SimpleAnticheat
}

func (w *WrappedAnticheatClient) AddPlayer(ctx context.Context, in *xyron.AddPlayerRequest, opts ...grpc.CallOption) (*xyron.PlayerReceipt, error) {
	return w.server.AddPlayer(ctx, in)
}

func (w *WrappedAnticheatClient) RemovePlayer(ctx context.Context, in *xyron.PlayerReceipt, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	return w.server.RemovePlayer(ctx, in)
}

func (w *WrappedAnticheatClient) Report(ctx context.Context, in *xyron.PlayerReport, opts ...grpc.CallOption) (*xyron.ReportResponse, error) {
	return w.server.Report(ctx, in)
}

func getField(s interface{}, n string) reflect.Value {
	return reflect.ValueOf(s).Elem().FieldByName(n)
}

func getUnexportedField(s interface{}, n string) interface{} {
	field := getField(s, n)
	return reflect.NewAt(field.Type(), unsafe.Pointer(field.UnsafeAddr())).Elem().Interface()
}

func getCurrentWorldTick(w *world.World) int64 {
	set := getUnexportedField(w, "set").(*world.Settings)
	set.Lock()
	defer set.Unlock()
	return set.CurrentTick
}

type BufferedDataQueue struct {
	mu *sync.Mutex
	m  map[int64][]*xyron.WildcardReportData
}

func NewBufferedDataQueue() *BufferedDataQueue {
	return &BufferedDataQueue{
		mu: &sync.Mutex{},
		m:  make(map[int64][]*xyron.WildcardReportData, 128),
	}
}

func (b *BufferedDataQueue) Add(tick int64, wdata *xyron.WildcardReportData) {
	b.mu.Lock()
	defer b.mu.Unlock()
	if _, ok := b.m[tick]; !ok {
		b.m[tick] = nil
	}
	b.m[tick] = append(b.m[tick], wdata)
}

func (b *BufferedDataQueue) Flush(ctx context.Context, c xyron.AnticheatClient, p *xyron.PlayerReceipt, tick int64) (*xyron.ReportResponse, error) {
	b.mu.Lock()
	var needSend []int64
	for k, _ := range b.m {
		if k <= tick {
			needSend = append(needSend, k)
		}
	}
	sorted := anticheat.ComparableSlice[int64](needSend)
	sorted.Sort()
	needSendMap := make(map[int64]*xyron.TimestampedReportData, len(needSend))
	for _, v := range sorted {
		needSendMap[v] = &xyron.TimestampedReportData{Data: b.m[v]}
		delete(b.m, v)
	}
	b.mu.Unlock()
	return c.Report(ctx, &xyron.PlayerReport{
		Player: p,
		Data:   needSendMap,
	})
}

func main() {
	log := logrus.New()
	log.Formatter = &logrus.TextFormatter{ForceColors: true}
	log.Level = logrus.DebugLevel

	c := &WrappedAnticheatClient{anticheat.NewSimpleAnticheatServer(implementation.Available)}

	chat.Global.Subscribe(chat.StdoutSubscriber{})

	conf, err := readConfig(log)
	if err != nil {
		log.Fatalln(err)
	}

	srv := conf.New()
	srv.CloseOnProgramEnd()

	srv.Listen()

	for srv.Accept(func(p *player.Player) {
		p.SetGameMode(world.GameModeSurvival)
		var hdrs []*xyron.WildcardReportData
		hdrs = append(hdrs, &xyron.WildcardReportData{Data: &xyron.WildcardReportData_GameModeData{
			GameModeData: &xyron.PlayerGameModeData{GameMode: ToXyronGameMode(p.GameMode())},
		}})
		handleEffects(p, &hdrs)
		go func() {
			pp, _ := c.AddPlayer(context.TODO(), &xyron.AddPlayerRequest{
				Player: &xyron.Player{
					//TODO Os
					Name: p.Name(),
				},
				Data: map[int64]*xyron.TimestampedReportData{0: {Data: hdrs}},
			})
			p.Handle(newPlayerHandler(log, p, pp, c))
		}()
	}) {
	}
}

func handleEffects(p *player.Player, hdrs *[]*xyron.WildcardReportData) {
	effects := getEffects(p)
	*hdrs = append(*hdrs, &xyron.WildcardReportData{Data: &xyron.WildcardReportData_EffectData{
		EffectData: &xyron.PlayerEffectData{Effect: effects},
	}})
}

func getEffects(p *player.Player) []*xyron.EffectFeature {
	var effects []*xyron.EffectFeature
	if e, ok := p.Effect(effect.Speed{}); ok {
		effects = append(effects, &xyron.EffectFeature{
			Amplifier: int32(e.Level()),
			IsSpeed:   true,
		})
	}
	if e, ok := p.Effect(effect.Haste{}); ok {
		effects = append(effects, &xyron.EffectFeature{
			Amplifier: int32(e.Level()),
			IsHaste:   true,
		})
	}
	if e, ok := p.Effect(effect.SlowFalling{}); ok {
		effects = append(effects, &xyron.EffectFeature{
			Amplifier:     int32(e.Level()),
			IsSlowFalling: true,
		})
	}
	if e, ok := p.Effect(effect.Levitation{}); ok {
		effects = append(effects, &xyron.EffectFeature{
			Amplifier:    int32(e.Level()),
			IsLevitation: true,
		})
	}
	if e, ok := p.Effect(effect.Slowness{}); ok {
		effects = append(effects, &xyron.EffectFeature{
			Amplifier:  int32(e.Level()),
			IsSlowness: true,
		})
	}
	if e, ok := p.Effect(effect.JumpBoost{}); ok {
		effects = append(effects, &xyron.EffectFeature{
			Amplifier:   int32(e.Level()),
			IsJumpBoost: true,
		})
	}
	return effects
}

type playerHandler struct {
	player.NopHandler
	log    *logrus.Logger
	p      *player.Player
	buf    *BufferedDataQueue
	pp     *xyron.PlayerReceipt
	c      xyron.AnticheatClient
	closed atomic.Bool
	ticker *time.Ticker
}

func newPlayerHandler(log *logrus.Logger, p *player.Player, pp *xyron.PlayerReceipt, c xyron.AnticheatClient) *playerHandler {
	hdr := &playerHandler{
		NopHandler: player.NopHandler{},
		p:          p,
		buf:        NewBufferedDataQueue(),
		pp:         pp,
		c:          c,
		closed:     atomic.Bool{},
	}
	hdr.ticker = time.NewTicker(time.Second / 5)
	hdr.closed.Store(false)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				hdr.log.Error(r)
			}
		}()
		for !hdr.closed.Load() {
			select {
			case _ = <-hdr.ticker.C:
				if !hdr.closed.Load() {
					if jdjm, err := hdr.buf.Flush(context.TODO(), hdr.c, hdr.pp, getCurrentWorldTick(hdr.p.World())); err != nil {
						hdr.p.Messagef("judgement error: %v", err)
					} else {
						for _, j := range jdjm.Judgements {
							hdr.p.Messagef("judgement: %v: %v message:%v", j.Type, j.Judgement.String(), j.Message)
						}
					}
				}
			}
		}
	}()
	return hdr
}
func (h *playerHandler) HandleTeleport(*event.Context, mgl64.Vec3) {}

func (h *playerHandler) HandleMove(_ *event.Context, newPos mgl64.Vec3, yaw, pitch float64) {
	h.buf.Add(getCurrentWorldTick(h.p.World()), &xyron.WildcardReportData{Data: &xyron.WildcardReportData_EffectData{
		EffectData: &xyron.PlayerEffectData{Effect: getEffects(h.p)},
	}})
	h.buf.Add(getCurrentWorldTick(h.p.World()), &xyron.WildcardReportData{Data: &xyron.WildcardReportData_MoveData{
		MoveData: &xyron.PlayerMoveData{
			NewPosition: h.getXyronPositionData(newPos, cube.Rotation{yaw, pitch}.Vec3()),
			Teleport:    true,
		},
	}})
}

func (h *playerHandler) HandleJump() {
	h.buf.Add(getCurrentWorldTick(h.p.World()), &xyron.WildcardReportData{Data: &xyron.WildcardReportData_ActionData{
		ActionData: &xyron.PlayerActionData{
			Position: h.getXyronPositionData(h.p.Position(), h.p.Rotation().Vec3()),
			Action:   xyron.PlayerAction_Jump,
		},
	}})
}

func (h *playerHandler) HandleQuit() {
	h.closed.Store(true)
	go h.c.RemovePlayer(context.TODO(), h.pp)
}

func getColliedBlocks(p *player.Player) []*xyron.BlockData {
	box := p.Type().BBox(p).Translate(p.Position().Add(mgl64.Vec3{0, -0.5000001, 0}))

	b := box.Grow(1)

	min, max := cube.PosFromVec3(b.Min()), cube.PosFromVec3(b.Max())

	var blocks []*xyron.BlockData

	for x := min[0]; x <= max[0]; x++ {
		for z := min[2]; z <= max[2]; z++ {
			for y := min[1]; y < max[1]; y++ {
				pos := cube.Pos{x, y, z}
				boxList := p.World().Block(pos).Model().BBox(pos, p.World())
				blk := p.World().Block(pos)
				for _, bb := range boxList {
					if bb.GrowVec3(mgl64.Vec3{0, 0.05}).Translate(pos.Vec3()).IntersectsWith(box) {
						bd, done := ToXyronBlockData(p.World(), blk, pos)
						if done {
							blocks = append(blocks, bd)
						}
					}
				}
			}
		}
	}
	return blocks
}

func getIntersectedBlocks(p *player.Player) []*xyron.BlockData {
	var blocks []*xyron.BlockData
	//TODO
	return blocks
}

func ToXyronBBox(bbox cube.BBox) *xyron.AxisAlignedBoundingBox {
	return &xyron.AxisAlignedBoundingBox{
		Min: ToXyronVec3(bbox.Min()),
		Max: ToXyronVec3(bbox.Max()),
	}
}

func ToXyronVec3(pos mgl64.Vec3) *xyron.Vec3F {
	return &xyron.Vec3F{
		X: float32(pos.X()),
		Y: float32(pos.Y()),
		Z: float32(pos.Z()),
	}
}

func ToXyronCubePos(pos cube.Pos) *xyron.Vec3I {
	return &xyron.Vec3I{
		X: int32(pos.X()),
		Y: int32(pos.Y()),
		Z: int32(pos.Z()),
	}
}

func ToXyronGameMode(g world.GameMode) xyron.GameMode {
	switch g {
	case world.GameModeSurvival:
		return xyron.GameMode_Survival
	case world.GameModeAdventure:
		return xyron.GameMode_Adventure
	case world.GameModeCreative:
		return xyron.GameMode_Creative
	case world.GameModeSpectator:
		return xyron.GameMode_Spectator
	default:
		panic(g)
	}
}

func ToXyronBlockData(w *world.World, blk world.Block, pos cube.Pos) (*xyron.BlockData, bool) {
	var bboxs []*xyron.AxisAlignedBoundingBox
	for _, bbox := range blk.Model().BBox(pos, w) {
		bboxs = append(bboxs, ToXyronBBox(bbox))
	}
	fric := float32(0.0)
	if f, ok := blk.(block.Frictional); ok {
		fric = float32(f.Friction())
	}
	_, solid := blk.Model().(model.Solid)
	_, air := blk.Model().(model.Empty)
	_, wtr := blk.(block.Water)
	_, lava := blk.(block.Lava)
	_, cl := blk.(block.Ladder)
	_, ice := blk.(block.PackedIce)
	_, ice2 := blk.(block.BlueIce)
	if air {
		return &xyron.BlockData{
			RelativePosition: ToXyronCubePos(pos),
			Feature: &xyron.BlockFeature{
				IsAir: true,
			},
		}, false
	}
	bd := &xyron.BlockData{
		RelativePosition: ToXyronCubePos(pos),
		Feature: &xyron.BlockFeature{
			CollisionBoxes: bboxs,
			Friction:       fric,
			IsSolid:        solid,
			IsLiquid:       wtr || lava,
			IsAir:          air,
			//dragonfly has no slime block
			IsSlime:     false,
			IsClimbable: cl,
			IsIce:       ice || ice2,
			//dragonfly has no cobweb
			IsCobweb: false,
			//dragonfly has no sweet berry
			IsSweetBerry: false,
		},
	}
	return bd, true
}

func (h *playerHandler) getXyronPositionData(pos, rot mgl64.Vec3) *xyron.EntityPositionData {
	xpos := cube.PosFromVec3(pos)
	xpos[1] = int(math.Floor(h.p.Type().BBox(h.p).Min()[1] - 0.50001))
	b, _ := ToXyronBlockData(h.p.World(), h.p.World().Block(xpos), xpos)
	return &xyron.EntityPositionData{
		Location: &xyron.Loc3F{
			Position:  ToXyronVec3(pos),
			Direction: ToXyronVec3(rot),
		},
		BoundingBox:             ToXyronBBox(h.p.Type().BBox(h.p)),
		BelowThatAffectMovement: b,
		IsImmobile:              h.p.Immobile(),
		IsOnGround:              h.p.OnGround(),
		AllowFlying:             h.p.GameMode().AllowsFlying(),
		IsFlying:                h.p.Flying(),
		HaveGravity:             true,
		CollidedBlocks:          getColliedBlocks(h.p),
		IntersectedBlocks:       getIntersectedBlocks(h.p),
	}
}

// readConfig reads the configuration from the config.toml file, or creates the
// file if it does not yet exist.
func readConfig(log server.Logger) (server.Config, error) {
	c := server.DefaultConfig()
	var zero server.Config
	if _, err := os.Stat("config.toml"); os.IsNotExist(err) {
		data, err := toml.Marshal(c)
		if err != nil {
			return zero, fmt.Errorf("encode default config: %v", err)
		}
		if err := os.WriteFile("config.toml", data, 0644); err != nil {
			return zero, fmt.Errorf("create default config: %v", err)
		}
		return c.Config(log)
	}
	data, err := os.ReadFile("config.toml")
	if err != nil {
		return zero, fmt.Errorf("read config: %v", err)
	}
	if err := toml.Unmarshal(data, &c); err != nil {
		return zero, fmt.Errorf("decode config: %v", err)
	}
	return c.Config(log)
}
