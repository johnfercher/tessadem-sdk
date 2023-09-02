package tessadem

import (
	"context"
	"encoding/json"
	"os"
)

type FileWriter interface {
	SaveArea(ctx context.Context, file string, response *AreaResponse) error
}

func NewFileWriter() FileWriter {
	return &fileWriter{}
}

type fileWriter struct {
}

func (f *fileWriter) SaveArea(ctx context.Context, file string, response *AreaResponse) error {
	osFile, err := os.Create(file)
	if err != nil {
		return err
	}
	defer osFile.Close()

	bytes, err := json.Marshal(response)
	if err != nil {
		return err
	}

	_, err = osFile.WriteString(string(bytes))
	if err != nil {
		return err
	}

	return nil
}
