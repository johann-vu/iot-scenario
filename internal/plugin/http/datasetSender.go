package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/johann-vu/iot-scenario/internal/domain"
)

type datasetSender struct {
	targetURL string
}

const resultEndpoint = "/results"

// Send implements domain.DatasetSender.
func (ds *datasetSender) Send(ctx context.Context, d domain.Dataset) error {
	datasetDTO := DatasetDTO{
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

func NewDatasetSender(targetURL string) domain.DatasetSender {
	return &datasetSender{targetURL: targetURL}
}
