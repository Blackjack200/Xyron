package anticheat

import (
	"github.com/blackjack200/xyron/xyron"
	"github.com/go-gl/mathgl/mgl64"
	"github.com/sirupsen/logrus"
	"time"
)

type InternalPlayer struct {
	log               *logrus.Logger
	lastReport        time.Time
	timestampThisTick int64
	checks            []any

	Name string

	Os       xyron.DeviceOS
	Input    xyron.InputMode
	GameMode xyron.GameMode

	Alive *BufferedTimestampedData[bool]

	effects []*xyron.EffectFeature

	Location      *BufferedData[*xyron.EntityPositionData]
	DeltaPosition *BufferedData[mgl64.Vec3]

	HeldItem *BufferedTimestampedData[*xyron.ItemData]

	Attack   *BufferedTimestampedData[*xyron.AttackData]
	Jump     *BufferedTimestampedData[float64]
	Eat      *BufferedTimestampedData[bool]
	Teleport *BufferedTimestampedData[mgl64.Vec3]

	Motion         *BufferedTimestampedData[mgl64.Vec3]
	MotionCoolDown int64

	OpenInventory  *BufferedTimestampedData[bool]
	CloseInventory *BufferedTimestampedData[bool]

	PlaceBlock *BufferedTimestampedData[*xyron.PlayerPlaceBlockData]
	BreakBlock *BufferedTimestampedData[*xyron.PlayerBreakBlockData]

	Sprinting *BufferedTimestampedData[bool]
	Sneaking  *BufferedTimestampedData[bool]
	Gliding   *BufferedTimestampedData[bool]
	Swimming  *BufferedTimestampedData[bool]
	Flying    *BufferedTimestampedData[bool]

	OnGround     *BufferedTimestampedData[bool]
	OnIce        *BufferedTimestampedData[bool]
	OnClimbable  *BufferedTimestampedData[bool]
	InCobweb     *BufferedTimestampedData[bool]
	InSweetBerry *BufferedTimestampedData[bool]

	IntersectedLiquid *BufferedTimestampedData[bool]
	IntersectedSolid  *BufferedTimestampedData[bool]

	InAirTick    uint32
	OnGroundTick uint32
	OnIceTick    uint32
}

func NewInternalPlayer(log *logrus.Logger, checks []any, os xyron.DeviceOS, name string) *InternalPlayer {
	return &InternalPlayer{
		log:               log,
		lastReport:        time.Now(),
		timestampThisTick: 0,
		checks:            checks,
		Name:              name,
		Os:                os,
		Input:             0,
		GameMode:          0,
		Alive:             NewBufferedTimestampedData(true),
		effects:           nil,
		Location:          NewBufferedData((*xyron.EntityPositionData)(nil)),
		DeltaPosition:     NewBufferedData(mgl64.Vec3{}),
		HeldItem:          NewBufferedTimestampedData((*xyron.ItemData)(nil)),
		Attack:            NewBufferedTimestampedData((*xyron.AttackData)(nil)),
		Jump:              NewBufferedTimestampedData(float64(0)),
		Eat:               NewBufferedTimestampedData(false),
		Teleport:          NewBufferedTimestampedData(mgl64.Vec3{}),
		Motion:            NewBufferedTimestampedData(mgl64.Vec3{}),
		MotionCoolDown:    0,
		OpenInventory:     NewBufferedTimestampedData(false),
		CloseInventory:    NewBufferedTimestampedData(false),
		PlaceBlock:        NewBufferedTimestampedData((*xyron.PlayerPlaceBlockData)(nil)),
		BreakBlock:        NewBufferedTimestampedData((*xyron.PlayerBreakBlockData)(nil)),
		Sprinting:         NewBufferedTimestampedData(false),
		Sneaking:          NewBufferedTimestampedData(false),
		Gliding:           NewBufferedTimestampedData(false),
		Swimming:          NewBufferedTimestampedData(false),
		Flying:            NewBufferedTimestampedData(false),
		OnGround:          NewBufferedTimestampedData(true),
		OnIce:             NewBufferedTimestampedData(false),
		OnClimbable:       NewBufferedTimestampedData(false),
		InCobweb:          NewBufferedTimestampedData(false),
		InSweetBerry:      NewBufferedTimestampedData(false),
		IntersectedLiquid: NewBufferedTimestampedData(false),
		IntersectedSolid:  NewBufferedTimestampedData(false),
		InAirTick:         0,
		OnGroundTick:      0,
		OnIceTick:         0,
	}
}

func (p *InternalPlayer) SetLocation(pos *xyron.EntityPositionData) {
	p.Location.Set(pos)
	if p.Location.Previous() != nil {
		prev := toVec3(p.Location.Previous().Position)
		cur := toVec3(pos.Position)
		p.DeltaPosition.Set(cur.Sub(prev))
	}
	if pos != nil {
		OnGround, OnIce, InCobweb, InSweetBerry, OnClimbable, IntersectedLiquid, IntersectedSolid := p.CheckGroundState(pos)
		p.OnGround.Set(p.timestampThisTick, OnGround)
		p.OnIce.Set(p.timestampThisTick, OnIce)
		p.InCobweb.Set(p.timestampThisTick, InCobweb)
		p.InSweetBerry.Set(p.timestampThisTick, InSweetBerry)
		p.OnClimbable.Set(p.timestampThisTick, OnClimbable)
		p.IntersectedLiquid.Set(p.timestampThisTick, IntersectedLiquid)
		p.IntersectedSolid.Set(p.timestampThisTick, IntersectedSolid)
	}
}

func (p *InternalPlayer) CheckGroundState(pos *xyron.EntityPositionData) (
	OnGround,
	OnIce,
	InCobweb,
	InSweetBerry,
	OnClimbable,
	IntersectedLiquid,
	IntersectedSolid bool,
) {
	check := func(checkFeature func(*xyron.BlockFeature) bool) func([]*xyron.BlockData) bool {
		return func(bb []*xyron.BlockData) bool {
			for _, b := range bb {
				if checkFeature(b.Feature) {
					return true
				}
			}
			return false
		}
	}
	checkSolid := check(func(f *xyron.BlockFeature) bool { return f.IsSolid })
	checkIce := check(func(f *xyron.BlockFeature) bool { return f.IsIce })
	checkCobweb := check(func(f *xyron.BlockFeature) bool { return f.IsCobweb })
	checkSweetBerry := check(func(f *xyron.BlockFeature) bool { return f.IsSweetBerry })
	checkClimbable := check(func(f *xyron.BlockFeature) bool { return f.IsClimbable })
	checkLiquid := check(func(f *xyron.BlockFeature) bool { return f.IsLiquid })

	OnGround = checkSolid(pos.CollidedBlocks) || checkSolid([]*xyron.BlockData{pos.BelowThatAffectMovement})
	OnIce = checkIce(pos.CollidedBlocks) || checkIce([]*xyron.BlockData{pos.BelowThatAffectMovement})
	InCobweb = checkCobweb(pos.CollidedBlocks) || checkCobweb(pos.IntersectedBlocks) || checkCobweb([]*xyron.BlockData{pos.BelowThatAffectMovement})
	InSweetBerry = checkSweetBerry(pos.CollidedBlocks) || checkSweetBerry(pos.IntersectedBlocks) || checkSweetBerry([]*xyron.BlockData{pos.BelowThatAffectMovement})
	OnClimbable = checkClimbable(pos.CollidedBlocks) || checkClimbable(pos.IntersectedBlocks) || checkClimbable([]*xyron.BlockData{pos.BelowThatAffectMovement})
	IntersectedLiquid = checkLiquid(pos.IntersectedBlocks) || checkLiquid([]*xyron.BlockData{pos.BelowThatAffectMovement})
	IntersectedSolid = checkSolid(pos.IntersectedBlocks)
	return
}

func (p *InternalPlayer) CurrentTimestamp() int64 {
	return p.timestampThisTick
}

func (p *InternalPlayer) Tick() {
	p.Motion.Set(p.timestampThisTick, mgl64.Vec3{})
	if !p.OnGround.Current().Get() {
		p.InAirTick++
		p.OnGroundTick = 0
	} else {
		p.InAirTick = 0
		p.OnGroundTick++
	}
	if p.OnIce.Current().Get() {
		p.OnIceTick++
	} else {
		p.OnIceTick = 0
	}
	if p.MotionCoolDown > 0 {
		p.MotionCoolDown--
	}
}

func (p *InternalPlayer) Effect(flag func(feature *xyron.EffectFeature) bool) float64 {
	for _, e := range p.effects {
		if flag(e) {
			return float64(e.Amplifier)
		}
	}
	return 0
}
