package main

import (
	"context"
	"fmt"
	"github.com/johnfercher/tessadem-sdk/pkg/tessadem"
)

func main() {
	ctx := context.TODO()
	apiKey := ""

	client := tessadem.NewClient(apiKey)
	writer := tessadem.NewFileWriter()

	request := &tessadem.AreaRequest{
		Units: tessadem.Meter,
		A: &tessadem.Vector2D{
			X: -22.956808892072303,
			Y: -43.14656384010024,
		},
		B: &tessadem.Vector2D{
			X: -22.936719026379752,
			Y: -43.18380019316494,
		},
	}

	response, err := client.GetArea(ctx, request)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	err = writer.SaveArea(ctx, "internal/data/file.json", response)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
