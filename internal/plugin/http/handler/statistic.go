package handler

import (
	"encoding/json"
	"net/http"
	"time"

	calculatestatistics "github.com/johann-vu/iot-scenario/internal/domain/calculateStatistics"
)

type statisticHandler struct {
	service calculatestatistics.Service
}

// ServeHTTP implements http.Handler.
func (sh *statisticHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	result, err := sh.service.Execute(r.Context(), time.Now().Add(-24*time.Hour), time.Now())
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func NewStatisticHandler(service calculatestatistics.Service) http.Handler {
	return &statisticHandler{service: service}
}
