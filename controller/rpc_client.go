package main

import (
	"context"
	"encoding/json"
	"log"
	"math"
	"strconv"
	"time"

	"github.com/torchiaf/Sensors/controller/config"
	"github.com/torchiaf/Sensors/rpc_client"
)

type Dht11 struct {
	T float64 `json:"t"`
	H float64 `json:"h"`
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {
	log.Printf("Config %+v", config.Config)

	client := rpc_client.New(context.Background())

	for {

		temperature := ""

		for _, module := range config.Config.Modules {
			log.Printf(" [x] Requesting on {%s, %s, %s}", module.Name, module.Type, module.RoutingKey)

			res, err := client.Read(
				module.RoutingKey,
				"dht11",
				[]string{},
			)
			failOnError(err, "Failed to handle RPC request: dht11")

			log.Printf(" [%s] [%s] Got %+v", module.Name, "dht11", res)

			if module.Name == "raspberrypi-0" {
				var obj Dht11
				err = json.Unmarshal([]byte(res), &obj)
				if err != nil {
					log.Fatalf("error: %v", err)
				}

				temperature = strconv.Itoa(int(math.Round(obj.T)))
			}

			if module.Name == "raspberrypi-1" {
				res1, err := client.Write(
					module.RoutingKey,
					"tm1637",
					[]string{
						"temperature",
						temperature,
					},
				)
				failOnError(err, "Failed to handle RPC request: tm1637")

				log.Printf(" [%s] [%s] Got %+v", module.Name, "tm1637", res1)
			}
		}

		time.Sleep(time.Second)
	}
}
