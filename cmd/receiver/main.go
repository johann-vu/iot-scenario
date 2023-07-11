package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	storedataset "github.com/johann-vu/iot-scenario/internal/domain/storeDataset"
	"github.com/johann-vu/iot-scenario/internal/plugin/http/handler"
	"github.com/johann-vu/iot-scenario/internal/plugin/storage/memory"
)

func main() {
	datasetRpository := memory.NewDatasetRepository()
	storeService := storedataset.NewService(datasetRpository)

	r := mux.NewRouter()

	r.Handle("/results", handler.NewDatasetHandler(storeService))

	fmt.Println(http.ListenAndServe(":8080", r))
}
