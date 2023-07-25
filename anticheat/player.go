package anticheat

import "github.com/blackjack200/xyron/xyron"

type InternalPlayer struct {
	checks        []any
	Os            xyron.DeviceOS
	Input         xyron.InputMode
	Name          string
	GameMode      xyron.GameMode
	effects       []*xyron.EffectFeature
	Motion        *BufferedData[*TimestampedData[*xyron.Vec3F]]
	Location      *BufferedData[*xyron.EntityPositionData]
	DeltaPosition *BufferedData[*xyron.Vec3F]
	Volatile      *TickedData[*VolatileData]

	Sprinting *BufferedData[bool]
	Sneaking  *BufferedData[bool]

	OnGround     *BufferedData[bool]
	OnIce        *BufferedData[bool]
	InCobweb     *BufferedData[bool]
	InSweetBerry *BufferedData[bool]
	OnClimbable  *BufferedData[bool]

	InAirTick        uint32
	OnGroundTick     uint32
	OnIceTick        uint32
	InCobwebTick     uint32
	InSweetBerryTick uint32
	currentTimestamp int64

	Teleport *BufferedData[int64]
}

func NewInternalPlayer(checks []any, os xyron.DeviceOS, name string) *InternalPlayer {
	return &InternalPlayer{
		checks:        checks,
		Os:            os,
		Name:          name,
		GameMode:      0,
		Motion:        NewBufferedData[*TimestampedData[*xyron.Vec3F]](nil),
		Location:      NewBufferedData[*xyron.EntityPositionData](nil),
		DeltaPosition: NewBufferedData[*xyron.Vec3F](nil),
		Volatile:      NewTickedData(&VolatileData{}),
		Sprinting:     NewBufferedData(false),
		Sneaking:      NewBufferedData(false),
		OnGround:      NewBufferedData(true),
		OnIce:         NewBufferedData(false),
		InCobweb:      NewBufferedData(false),
		InSweetBerry:  NewBufferedData(false),
		OnClimbable:   NewBufferedData(false),
		Teleport:      NewBufferedData[int64](0),
	}
}

func (p *InternalPlayer) GetVolatile() *VolatileData {
	return p.Volatile.Get()
}

func (p *InternalPlayer) SetLocation(pos *xyron.EntityPositionData) {
	p.Location.Set(pos)
	if p.Location.Previous() != nil {
		prev := p.Location.Previous().Location.Position
		cur := pos.Location.Position
		p.DeltaPosition.Set(&xyron.Vec3F{
			X: cur.X - prev.X,
			Y: cur.Y - prev.Y,
			Z: cur.Z - prev.Z,
		})
	}
	if pos != nil {
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
		checkSolid := check(func(f *xyron.BlockFeature) bool {
			return f.IsSolid
		})
		checkIce := check(func(f *xyron.BlockFeature) bool {
			return f.IsIce
		})
		checkCobweb := check(func(f *xyron.BlockFeature) bool {
			return f.IsCobweb
		})
		checkSweetBerry := check(func(f *xyron.BlockFeature) bool {
			return f.IsSweetBerry
		})
		checkClimbable := check(func(f *xyron.BlockFeature) bool {
			return f.IsClimbable
		})
		p.OnGround.Set(checkSolid(pos.CollidedBlocks) || checkSolid(pos.IntersectedBlocks))
		p.OnIce.Set(checkIce(pos.CollidedBlocks) || checkIce(pos.IntersectedBlocks))
		p.InCobweb.Set(checkCobweb(pos.CollidedBlocks) || checkCobweb(pos.IntersectedBlocks))
		p.InSweetBerry.Set(checkSweetBerry(pos.CollidedBlocks) || checkSweetBerry(pos.IntersectedBlocks))
		p.OnClimbable.Set(checkClimbable(pos.CollidedBlocks) || checkClimbable(pos.IntersectedBlocks))
	}
}

type VolatileData struct {
	Jumped     bool
	Teleported bool
}

func (p *InternalPlayer) CurrentTimestamp() int64 {
	return p.currentTimestamp
}

func (p *InternalPlayer) Tick() {
	p.Volatile.Set(&VolatileData{})
	p.Motion.Set(nil)
	if !p.OnGround.Current() {
		p.InAirTick++
		p.OnGroundTick = 0
	} else {
		p.InAirTick = 0
		p.OnGroundTick++
	}
	if p.OnIce.Current() {
		p.OnIceTick++
	} else {
		p.OnIceTick = 0
	}
	if p.InCobweb.Current() {
		p.InCobwebTick++
	} else {
		p.InCobwebTick = 0
	}
	if p.InSweetBerry.Current() {
		p.InSweetBerryTick++
	} else {
		p.InSweetBerryTick = 0
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
