package main

import (
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
	useMemory        bool
	connectionString string
	port             int
)

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
	r.Handle("/statistics", http.NewStatisticHandler(statisticService))

	fmt.Println(gohttp.ListenAndServe(":8080", r))
}

func loadConfig() {

	flag.StringVar(&connectionString, "connectionString", "root:mysecretpassword@tcp(localhost:3306)/mydb?parseTime=true", "Connection String to connect to MongoDB")
	flag.BoolVar(&useMemory, "useMemory", false, "Whether to store data in memory")
	flag.IntVar(&port, "port", 8080, "The port the receiver is listening on")

	flag.Parse()

	log.Println("Config has been loaded:")
	log.Printf("Port: \t%d", port)
	if !useMemory {
		log.Println("Database: \tSQL")
		log.Printf("Connection String: \t%v...", connectionString[:7])
	} else {
		log.Println("Database: \tMemory")
	}
}
