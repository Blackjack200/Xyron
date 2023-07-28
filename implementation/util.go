package implementation

import (
	"github.com/blackjack200/xyron/xyron"
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/go-gl/mathgl/mgl64"
	"math"
)

func toVec3(pos *xyron.Vec3F) mgl64.Vec3 {
	return mgl64.Vec3{
		float64(pos.X),
		float64(pos.Y),
		float64(pos.Z),
	}
}

func ToRotation(vec3 mgl64.Vec3) cube.Rotation {
	pitchRad := math.Asin(-vec3.Y())
	m := math.Cos(pitchRad)
	yawRad := math.Acos(vec3.Z() / m)
	return cube.Rotation{
		mgl64.RadToDeg(yawRad),
		mgl64.RadToDeg(pitchRad),
	}
}

func isZero(y float64) bool {
	return math.Abs(mgl64.Clamp(y, -epsilon, epsilon)) > epsilon
}
