package sender

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/johann-vu/iot-scenario/internal/domain"
	"github.com/johann-vu/iot-scenario/internal/plugin/http/dto"
)

type dataset struct {
	targetURL string
}

const resultEndpoint = "/results"

// Send implements domain.DatasetSender.
func (ds *dataset) Send(ctx context.Context, d domain.Dataset) error {
	datasetDTO := dto.DatasetDTO{
		SensorID:      d.SensorID,
		UnixTimestamp: d.Timestamp.Unix(),
		Value:         d.Value,
	}

	body := bytes.NewBuffer([]byte{})

	err := json.NewEncoder(body).Encode(datasetDTO)
	if err != nil {
		return err
	}

	resp, err := http.Post(ds.targetURL+resultEndpoint, "application/json", body)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode == 201 {
		return nil
	}

	return fmt.Errorf("unexpected status code in response: %d (%s)", resp.StatusCode, http.StatusText(resp.StatusCode))
}

func NewDataset(targetURL string) domain.DatasetSender {
	return &dataset{targetURL: targetURL}
}
