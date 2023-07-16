package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/google/uuid"
	senddataset "github.com/johann-vu/iot-scenario/internal/domain/sendDataset"
	"github.com/johann-vu/iot-scenario/internal/plugin/http/sender"
)

var (
	maxValue    float64
	minValue    float64
	receiverURL string
	sensorID    string
)

func main() {
	loadConfig()

	generator := senddataset.NewRandomGenerator(maxValue, minValue)
	sender := sender.NewDataset(receiverURL)
	service := senddataset.NewService(sensorID, generator, sender)

	for {
		service.Execute(context.Background())
		time.Sleep(5 * time.Second)
	}
}

func loadConfig() {

	flag.Float64Var(&maxValue, "max", 250, "maximum value the sensor can report")
	flag.Float64Var(&minValue, "min", -250, "minimum value the sensor can report")
	flag.StringVar(&receiverURL, "url", "http://localhost:8080", "url of the receiver")
	flag.StringVar(&sensorID, "id", uuid.New().String()[:4], "id of the sensor")

	flag.Parse()

	log.Println("Config has been loaded:")
	log.Printf("Maximum: \t%v", maxValue)
	log.Printf("Minimum: \t%v", minValue)
	log.Printf("Receiver: \t%s", receiverURL)
	log.Printf("ID: \t%s", sensorID)

}
