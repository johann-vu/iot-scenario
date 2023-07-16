package http

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator"
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

	var d DatasetDTO
	err := json.NewDecoder(r.Body).Decode(&d)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = dh.validate.StructCtx(r.Context(), d)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = dh.storeService.Execute(r.Context(), d.ToDomain())
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func NewDatasetHandler(service storedataset.Service) http.Handler {
	return &datasetHandler{storeService: service, validate: validator.New()}
}
