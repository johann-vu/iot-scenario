package http

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-playground/validator"
	"github.com/johann-vu/iot-scenario/internal/domain"
	storedataset "github.com/johann-vu/iot-scenario/internal/domain/storeDataset"
)

type datasetHandler struct {
	storeService storedataset.Service
	validate     *validator.Validate
}

// ServeHTTP implements http.Handler.
func (dh *datasetHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if r.Header.Get("content-type") != "application/json" {
		http.Error(w, http.StatusText(http.StatusUnsupportedMediaType), http.StatusUnsupportedMediaType)
		return
	}

	var dto DatasetDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = dh.validate.StructCtx(r.Context(), dto)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = dh.storeService.Execute(r.Context(), dto.ToDomain())
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func NewDatasetHandler(service storedataset.Service) http.Handler {
	return &datasetHandler{storeService: service, validate: validator.New()}
}

type DatasetDTO struct {
	SensorID      string  `json:"sensorId" validate:"required"`
	UnixTimestamp int64   `json:"unixTimestamp" validate:"required,gte=0"`
	Value         float32 `json:"value" validate:"required"`
}

func (d DatasetDTO) ToDomain() domain.Dataset {
	return domain.Dataset{
		SensorID:  d.SensorID,
		Timestamp: time.Unix(d.UnixTimestamp, 0),
		Value:     d.Value,
	}
}
