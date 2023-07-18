package http

import (
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"time"

	calculatestatistics "github.com/johann-vu/iot-scenario/internal/domain/calculateStatistics"
)

type statisticHandler struct {
	service           calculatestatistics.Service
	dashboardTemplate *template.Template
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

	if strings.HasPrefix(r.Header.Get("accept"), "text/html") {
		result.Maximum.Timestamp = result.Maximum.Timestamp.Local()
		result.Minimum.Timestamp = result.Minimum.Timestamp.Local()
		for i := 0; i < len(result.Recent); i++ {
			result.Recent[i].Timestamp = result.Recent[i].Timestamp.Local()
		}
		sh.dashboardTemplate.Execute(w, result)
		return
	}

	resultDto := StatisticsDTO{
		Average:           result.Average,
		StandardDeviation: result.StandardDeviation,
		Minimum: DatasetDTO{
			SensorID:      result.Minimum.SensorID,
			UnixTimestamp: result.Minimum.Timestamp.Unix(),
			Value:         result.Minimum.Value,
		},
		Maximum: DatasetDTO{
			SensorID:      result.Maximum.SensorID,
			UnixTimestamp: result.Maximum.Timestamp.Unix(),
			Value:         result.Maximum.Value,
		},
		Count:  result.Count,
		Slope:  result.LinearRegression.Slope,
		Recent: make([]DatasetDTO, len(result.Recent)),
	}

	for i := 0; i < len(result.Recent); i++ {
		resultDto.Recent[i] = DatasetDTO{
			SensorID:      result.Recent[i].SensorID,
			UnixTimestamp: result.Recent[i].Timestamp.Unix(),
			Value:         result.Recent[i].Value,
		}
	}

	json.NewEncoder(w).Encode(resultDto)
}

func NewStatisticHandler(service calculatestatistics.Service, dashboardFile []byte) http.Handler {

	t, err := template.New("dashboard").Parse(string(dashboardFile))
	if err != nil {
		panic(err)
	}

	return &statisticHandler{service: service, dashboardTemplate: t}
}
