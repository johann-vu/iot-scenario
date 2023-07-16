package http

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	calculatestatistics "github.com/johann-vu/iot-scenario/internal/domain/calculateStatistics"
)

type statisticHandler struct {
	service calculatestatistics.Service
}

// ServeHTTP implements http.Handler.
func (sh *statisticHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	timespan := -60 * time.Minute

	if timespanQuery := r.URL.Query().Get("timespan"); timespanQuery != "" {
		parsed, err := strconv.Atoi(timespanQuery)
		if err == nil {
			timespan = time.Duration(-1*parsed) * time.Minute
		}
	}

	result, err := sh.service.Execute(r.Context(), time.Now().Add(timespan), time.Now())
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func NewStatisticHandler(service calculatestatistics.Service) http.Handler {
	return &statisticHandler{service: service}
}
