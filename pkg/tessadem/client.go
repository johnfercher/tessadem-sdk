package tessadem

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client interface {
	GetArea(ctx context.Context, areaRequest *AreaRequest) (*AreaResponse, error)
}

func NewClient(apiKey string) Client {
	return &client{
		apiKey: apiKey,
	}
}

type client struct {
	apiKey string
}

func (c *client) GetArea(ctx context.Context, areaRequest *AreaRequest) (*AreaResponse, error) {
	//areaRequest.Square()
	proportions := areaRequest.GetProportions()
	topX := areaRequest.Northeast().X
	topY := areaRequest.Northeast().Y
	bottomX := areaRequest.Southwest().X
	bottomY := areaRequest.Southwest().Y
	rows := proportions.X
	cols := proportions.Y
	unit := areaRequest.Units
	mode := Area

	apiUrl := fmt.Sprintf("https://tessadem.com/api/elevation?key=%s&locations=%f,%f|%f,%f&mode=%s&units=%s&rows=%d&columns=%d", c.apiKey, topX, topY, bottomX, bottomY, mode, unit, int(rows), int(cols))

	request, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-Type", "application/json; charset=utf-8")

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	areaResponse := &AreaResponse{}
	err = json.Unmarshal(responseBody, areaResponse)
	if err != nil {
		return nil, err
	}

	return areaResponse, nil
}
