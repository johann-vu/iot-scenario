package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/google/uuid"
	senddataset "github.com/johann-vu/iot-scenario/internal/domain/sendDataset"
	"github.com/johann-vu/iot-scenario/internal/plugin/http"
)

var (
	maxValue    float64
	minValue    float64
	receiverURL string
	sensorID    string
	interval    int
	random      bool
	wave        bool
)

func main() {
	loadConfig()

	generator := senddataset.NewWaveGenerator(maxValue, minValue, 0.1)
	if random {
		generator = senddataset.NewRandomGenerator(maxValue, minValue)
	}
	sender := http.NewDatasetSender(receiverURL)
	service := senddataset.NewService(sensorID, generator, sender)

	for {
		service.Execute(context.Background())
		time.Sleep(time.Duration(interval) * time.Second)
	}
}

func loadConfig() {

	flag.Float64Var(&maxValue, "max", 100, "maximum value the sensor can report")
	flag.Float64Var(&minValue, "min", 0, "minimum value the sensor can report")
	flag.StringVar(&receiverURL, "url", "http://localhost:8080", "url of the receiver")
	flag.StringVar(&sensorID, "id", uuid.New().String()[:4], "id of the sensor")
	flag.IntVar(&interval, "interval", 3, "interval between sending data in seconds")
	flag.BoolVar(&random, "random", false, "send random values")
	flag.BoolVar(&wave, "wave", true, "send values following a wave function")

	flag.Parse()

	log.Println("Config has been loaded:")
	log.Printf("Maximum: \t%v", maxValue)
	log.Printf("Minimum: \t%v", minValue)
	log.Printf("Receiver: \t%s", receiverURL)
	log.Printf("ID: \t%s", sensorID)
	log.Printf("Interval: \t%d", interval)
	if random {
		log.Println("Mode: \tRandom")
	} else {
		log.Println("Mode: \tWave")
	}
}
