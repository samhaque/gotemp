package main

import (
	"encoding/json"
	"fmt"
	"github.com/d2r2/go-dht"
	"log"
)

type Payload struct {
	Temperature float32 `json:"temperature"`
	Humidity    float32 `json:"humidity"`
	Retries     int     `json:"retries"`
}

func main() {
	for {
		temperature, humidity, retried, err :=
			dht.ReadDHTxxWithRetry(dht.DHT11, 7, false, 10)
		if err != nil {
			log.Println("Failed to collect sensor data: ", err)
		}
		payload := &Payload{
			Temperature: temperature,
			Humidity:    humidity,
			Retries:     retried,
		}
		payloadJson, err := json.Marshal(payload)
		if err != nil {
			log.Println("Failed to marshal payload: ", err)
		}
		fmt.Println(string(payloadJson))
	}
}
