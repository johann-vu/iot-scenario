package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	generatedataset "github.com/johann-vu/iot-scenario/internal/domain/generateDataset"
	"github.com/johann-vu/iot-scenario/internal/plugin/http/dto"
)

func main() {
	generator := generatedataset.NewService("e", 0, 100)
	for i := 0; i < 30; i++ {
		time.Sleep(time.Second * 3)

		dataset, _ := generator.Execute()

		w := bytes.NewBuffer([]byte{})
		json.NewEncoder(w).Encode(dto.DatasetDTO{
			SensorID:      dataset.SensorID,
			Value:         dataset.Value,
			UnixTimestamp: dataset.Timestamp.Unix(),
		})

		http.Post("http://localhost:8080/results", "application/json", w)
	}
}
