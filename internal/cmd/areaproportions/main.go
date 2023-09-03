package main

import (
	"fmt"
	"github.com/johnfercher/tessadem-sdk/pkg/tessadem"
)

func main() {
	request := &tessadem.AreaRequest{
		Units: tessadem.Meter,
		Northeast: &tessadem.Vector2D{
			X: -22.510677151874123,
			Y: -43.18595653686086,
		},
		Southwest: &tessadem.Vector2D{
			X: -22.503520257853395,
			Y: -43.170512111338226,
		},
	}

	proportions := request.GetProportions()
	fmt.Println(proportions)
}
