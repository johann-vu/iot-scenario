package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	gohttp "net/http"

	"github.com/gorilla/mux"
	calculatestatistics "github.com/johann-vu/iot-scenario/internal/domain/calculateStatistics"
	storedataset "github.com/johann-vu/iot-scenario/internal/domain/storeDataset"
	"github.com/johann-vu/iot-scenario/internal/plugin/http"
	"github.com/johann-vu/iot-scenario/internal/plugin/storage"
)

var (
	connectionString string
	port             int
)

//go:embed index.html
var dashboardFile []byte

func main() {

	loadConfig()

	datasetRpository, err := storage.NewSQLDatasetRepository(connectionString)
	if err != nil {
		log.Panicf("connecting to database: %v", err)
	}

	storeService := storedataset.NewService(datasetRpository)
	statisticService := calculatestatistics.NewService(datasetRpository)

	r := mux.NewRouter()

	r.Handle("/results", http.NewDatasetHandler(storeService))
	r.Handle("/statistics", http.NewStatisticHandler(statisticService, dashboardFile))

	fmt.Println(gohttp.ListenAndServe(":8080", r))
}

func loadConfig() {

	flag.StringVar(&connectionString, "connectionString", "", "Connection String to connect to MySQL")
	flag.IntVar(&port, "port", 8080, "The port the receiver is listening on")

	flag.Parse()

	log.Println("Config has been loaded:")
	log.Printf("Port: \t%d", port)
	log.Println("Database: \tSQL")
	log.Printf("Connection String: \t%v...", connectionString[:7])
}
