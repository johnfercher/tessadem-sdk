package tessadem

import (
	"context"
	"encoding/json"
	"os"
)

type FileReader interface {
	ReadArea(ctx context.Context, file string) (*AreaResponse, error)
}

func NewFileReader() FileReader {
	return &fileReader{}
}

type fileReader struct {
}

func (f *fileReader) ReadArea(ctx context.Context, file string) (*AreaResponse, error) {
	bytes, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}

	response := &AreaResponse{}
	err = json.Unmarshal(bytes, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
