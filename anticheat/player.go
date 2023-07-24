package anticheat

import "github.com/blackjack200/xyron/xyron"

type InternalPlayer struct {
	Os       xyron.DeviceOS
	Input    xyron.InputMode
	Name     string
	GameMode xyron.GameMode
	Effects  map[string]*xyron.EffectFeature
	Motion   *BufferedData[*xyron.Vec3F]
	Location *BufferedData[*xyron.EntityPositionData]
	Volatile *TickedData[*VolatileData]

	OnGround     *BufferedData[bool]
	InAirTick    uint32
	OnGroundTick uint32
}

func NewInternalPlayer(os xyron.DeviceOS, name string) *InternalPlayer {
	return &InternalPlayer{
		Os:       os,
		Name:     name,
		GameMode: 0,
		Effects:  make(map[string]*xyron.EffectFeature),
		Motion:   NewBufferedData[*xyron.Vec3F](nil),
		Location: NewBufferedData[*xyron.EntityPositionData](nil),
		Volatile: NewTickedData(&VolatileData{}),
		OnGround: NewBufferedData(true),
	}
}

func (p *InternalPlayer) GetVolatile() *VolatileData {
	return p.Volatile.Get()
}

func (p *InternalPlayer) SetLocation(pos *xyron.EntityPositionData) {
	p.Location.Set(pos)
	if pos != nil {
		check := func(bb []*xyron.BlockData) bool {
			for _, b := range bb {
				if b.Feature.IsSolid {
					return true
				}
			}
			return false
		}
		p.OnGround.Set(check(pos.CollidedBlocks) || check(pos.IntersectedBlocks))
	}
}

type VolatileData struct {
	Jumped     bool
	Teleported bool
}

func (p *InternalPlayer) Tick() {
	p.Volatile.Reset()
	p.Motion.Set(nil)
	if !p.OnGround.Current() {
		p.InAirTick++
		p.OnGroundTick = 0
	} else {
		p.InAirTick = 0
		p.OnGroundTick++
	}
}

func (p *InternalPlayer) HasEffect(flag func(feature *xyron.EffectFeature) bool) (*xyron.EffectFeature, bool) {
	for _, e := range p.Effects {
		if flag(e) {
			return e, true
		}
	}
	return nil, false
}
