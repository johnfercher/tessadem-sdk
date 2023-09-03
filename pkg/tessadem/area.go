package tessadem

import (
	"math"
)

const (
	MaxArea float64 = 16384
)

type AreaRequest struct {
	Units     Unit
	Northeast *Vector2D
	Southwest *Vector2D
}

func (a *AreaRequest) GetProportions() *Vector2D {
	xRadians := math.Abs(a.Northeast.X - a.Southwest.X)
	yRadians := math.Abs(a.Northeast.Y - a.Southwest.Y)

	xDegrees := RadiansToDegrees(xRadians)
	yDegrees := RadiansToDegrees(yRadians)

	xMeters := DegreesToMeters(xDegrees)
	yMeters := DegreesToMeters(yDegrees)

	return NormalizeVector2DInMaxValue(&Vector2D{xMeters, yMeters})
}

type AreaResponse struct {
	Results [][]*Vector3D `json:"results"`
}
