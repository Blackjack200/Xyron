package anticheat

import (
	"github.com/blackjack200/xyron/xyron"
	"github.com/go-gl/mathgl/mgl64"
	"github.com/sirupsen/logrus"
	"sync"
)

type InternalPlayer struct {
	tickHandlingMu *sync.Mutex
	log            *logrus.Logger
	checks         []any

	Os            xyron.DeviceOS
	Input         xyron.InputMode
	Name          string
	GameMode      xyron.GameMode
	Alive         *BufferedTimestampedData[bool]
	effects       []*xyron.EffectFeature
	Motion        *BufferedTimestampedData[mgl64.Vec3]
	Location      *BufferedData[*xyron.EntityPositionData]
	DeltaPosition *BufferedData[mgl64.Vec3]
	Volatile      *TickedData[*VolatileData]

	Sprinting      *BufferedTimestampedData[bool]
	Sneaking       *BufferedTimestampedData[bool]
	Gliding        *BufferedTimestampedData[bool]
	Swimming       *BufferedTimestampedData[bool]
	Flying         *BufferedTimestampedData[bool]
	OpenInventory  *BufferedTimestampedData[bool]
	CloseInventory *BufferedTimestampedData[bool]
	Attack         *BufferedTimestampedData[*xyron.AttackData]

	OnGround          *BufferedTimestampedData[bool]
	OnIce             *BufferedTimestampedData[bool]
	InCobweb          *BufferedTimestampedData[bool]
	IntersectedLiquid *BufferedTimestampedData[bool]
	InSweetBerry      *BufferedTimestampedData[bool]
	OnClimbable       *BufferedTimestampedData[bool]

	InAirTick         uint32
	OnGroundTick      uint32
	OnIceTick         uint32
	timestampThisTick int64

	Teleport *BufferedTimestampedData[mgl64.Vec3]
}

func NewInternalPlayer(log *logrus.Logger, checks []any, os xyron.DeviceOS, name string) *InternalPlayer {
	return &InternalPlayer{
		tickHandlingMu:    &sync.Mutex{},
		log:               log,
		checks:            checks,
		Os:                os,
		Input:             0,
		Name:              name,
		GameMode:          0,
		Alive:             NewBufferedTimestampedData(true),
		effects:           nil,
		Motion:            NewBufferedTimestampedData(mgl64.Vec3{}),
		Location:          NewBufferedData[*xyron.EntityPositionData](nil),
		DeltaPosition:     NewBufferedData[mgl64.Vec3](mgl64.Vec3{}),
		Volatile:          NewTickedData(&VolatileData{}),
		Sprinting:         NewBufferedTimestampedData(false),
		Sneaking:          NewBufferedTimestampedData(false),
		Gliding:           NewBufferedTimestampedData(false),
		Swimming:          NewBufferedTimestampedData(false),
		Flying:            NewBufferedTimestampedData(false),
		OpenInventory:     NewBufferedTimestampedData(false),
		CloseInventory:    NewBufferedTimestampedData(false),
		Attack:            NewBufferedTimestampedData[*xyron.AttackData](nil),
		OnGround:          NewBufferedTimestampedData(true),
		OnIce:             NewBufferedTimestampedData(false),
		InCobweb:          NewBufferedTimestampedData(false),
		IntersectedLiquid: NewBufferedTimestampedData(false),
		InSweetBerry:      NewBufferedTimestampedData(false),
		OnClimbable:       NewBufferedTimestampedData(false),
		InAirTick:         0,
		OnGroundTick:      0,
		OnIceTick:         0,
		timestampThisTick: 0,
		Teleport:          NewBufferedTimestampedData(mgl64.Vec3{}),
	}
}

func (p *InternalPlayer) GetVolatile() *VolatileData {
	return p.Volatile.Get()
}

func (p *InternalPlayer) SetLocation(pos *xyron.EntityPositionData) {
	p.Location.Set(pos)
	if p.Location.Previous() != nil {
		prev := toVec3(p.Location.Previous().Position)
		cur := toVec3(pos.Position)
		p.DeltaPosition.Set(cur.Sub(prev))
	}
	if pos != nil {
		OnGround, OnIce, InCobweb, InSweetBerry, OnClimbable, IntersectedLiquid := p.CheckGroundState(pos)
		p.OnGround.Set(p.timestampThisTick, OnGround)
		p.OnIce.Set(p.timestampThisTick, OnIce)
		p.InCobweb.Set(p.timestampThisTick, InCobweb)
		p.InSweetBerry.Set(p.timestampThisTick, InSweetBerry)
		p.OnClimbable.Set(p.timestampThisTick, OnClimbable)
		p.IntersectedLiquid.Set(p.timestampThisTick, IntersectedLiquid)
	}
}

func (p *InternalPlayer) CheckGroundState(pos *xyron.EntityPositionData) (
	OnGround, OnIce, InCobweb, InSweetBerry, OnClimbable bool, IntersectedLiquid bool,
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

	OnGround = checkSolid(pos.CollidedBlocks) || checkSolid(pos.IntersectedBlocks) || checkSolid([]*xyron.BlockData{pos.BelowThatAffectMovement})
	OnIce = checkIce(pos.CollidedBlocks) || checkIce(pos.IntersectedBlocks) || checkIce([]*xyron.BlockData{pos.BelowThatAffectMovement})
	InCobweb = checkCobweb(pos.CollidedBlocks) || checkCobweb(pos.IntersectedBlocks) || checkCobweb([]*xyron.BlockData{pos.BelowThatAffectMovement})
	InSweetBerry = checkSweetBerry(pos.CollidedBlocks) || checkSweetBerry(pos.IntersectedBlocks) || checkSweetBerry([]*xyron.BlockData{pos.BelowThatAffectMovement})
	OnClimbable = checkClimbable(pos.CollidedBlocks) || checkClimbable(pos.IntersectedBlocks) || checkClimbable([]*xyron.BlockData{pos.BelowThatAffectMovement})
	IntersectedLiquid = checkLiquid(pos.IntersectedBlocks) || checkLiquid([]*xyron.BlockData{pos.BelowThatAffectMovement})
	return
}

type VolatileData struct {
	Jumped     bool
	Teleported bool
}

func (p *InternalPlayer) CurrentTimestamp() int64 {
	return p.timestampThisTick
}

func (p *InternalPlayer) Tick() {
	p.Volatile.Set(&VolatileData{})
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
}

func (p *InternalPlayer) GetEffect(flag func(feature *xyron.EffectFeature) bool) (*xyron.EffectFeature, bool) {
	for _, e := range p.effects {
		if flag(e) {
			return e, true
		}
	}
	return nil, false
}

func (p *InternalPlayer) HasEffect(flag func(feature *xyron.EffectFeature) bool) bool {
	for _, e := range p.effects {
		if flag(e) {
			return true
		}
	}
	return false
}
