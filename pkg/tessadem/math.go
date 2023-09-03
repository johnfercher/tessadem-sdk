package tessadem

import (
	"fmt"
	"math"
)

func RadiansToDegrees(radians float64) float64 {
	return radians * (180.0 / math.Pi)
}

func DegreesToMeters(degrees float64) float64 {
	// https://www.sco.wisc.edu/2022/01/21/how-big-is-a-degree/#:~:text=Therefore%20we%20can%20easily%20compute,further%20subdivisions%20of%20a%20degree.&text=circumference%20of%2025%2C000%20miles.
	return degrees * 111000.0
}

func NormalizeVector2DInMaxValue(v *Vector2D) *Vector2D {
	x := v.X
	y := v.Y

	for {
		if x*y < MaxArea {
			break
		}

		x, y = DecreaseHalf(x, y)
	}

	current := 0
	for {
		newX, newY := IncreasePercent(x, y, PercentProgression[current])
		if newX*newY > MaxArea {
			if current < len(PercentProgression)-1 {
				current++
				continue
			} else {
				break
			}
		}

		x, y = newX, newY
	}

	return &Vector2D{math.Round(x), math.Round(y)}
}

func DecreaseHalf(x, y float64) (float64, float64) {
	return x / 2.0, y / 2.0
}

var PercentProgression = []float64{1.75, 1.5, 1.25, 1.10, 1.05, 1.01}

func IncreasePercent(x, y, percent float64) (float64, float64) {
	fmt.Println(x, y, percent)
	return x * percent, y * percent
}
