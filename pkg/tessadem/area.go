package tessadem

import (
	"math"
)

const (
	MaxArea float64 = 16384
)

type AreaRequest struct {
	Units Unit
	A     *Vector2D
	B     *Vector2D
}

func (a *AreaRequest) Square() {
	absX := math.Abs(a.Northeast().X - a.Southwest().X)
	absY := math.Abs(a.Northeast().Y - a.Southwest().Y)

	max := absX
	if absY > absX {
		max = absY
	}

	minX := a.Northeast().X
	if a.Southwest().X < minX {
		minX = a.Southwest().X
	}

	minY := a.Northeast().Y
	if a.Southwest().Y < minY {
		minY = a.Southwest().Y
	}

	a.A.X = minX
	a.A.Y = minY

	a.B.X = minX + max
	a.B.Y = minY + max
}

func (a *AreaRequest) Northeast() *Vector2D {
	topX := a.A.X
	if a.B.X < topX {
		topX = a.B.X
	}

	topY := a.A.Y
	if a.B.Y < topY {
		topY = a.B.Y
	}

	return &Vector2D{
		X: topX,
		Y: topY,
	}
}

func (a *AreaRequest) Southwest() *Vector2D {
	bottomX := a.A.X
	if a.B.X > bottomX {
		bottomX = a.B.X
	}

	bottomY := a.A.Y
	if a.B.Y > bottomY {
		bottomY = a.B.Y
	}

	return &Vector2D{
		X: bottomX,
		Y: bottomY,
	}
}

func (a *AreaRequest) GetProportions() *Vector2D {
	xRadians := math.Abs(a.Northeast().X - a.Southwest().X)
	yRadians := math.Abs(a.Northeast().Y - a.Southwest().Y)

	xDegrees := RadiansToDegrees(xRadians)
	yDegrees := RadiansToDegrees(yRadians)

	xMeters := DegreesToMeters(xDegrees)
	yMeters := DegreesToMeters(yDegrees)

	return NormalizeVector2DInMaxValue(&Vector2D{xMeters, yMeters})
}

type AreaResponse struct {
	Results [][]*Vector3D `json:"results"`
}
