package main

import (
	"context"
	"fmt"

	rpc_client "github.com/torchiaf/Sensors/rpc_client"
)

type Message struct {
	// TODO refactoring
	Device string   `json:"device"`
	Action string   `json:"action"`
	Args   []string `json:"args"`
}

// func toSelector(labels *metav1.LabelSelector) (labels.Selector, error) {
// 	return metav1.LabelSelectorAsSelector(labels)
// }

func main() {

	// matchLabels := make(map[string]string)

	// matchLabels["fleet.cattle.io/benchmark"] = "true"

	// var vvv *metav1.LabelSelector = &metav1.LabelSelector{
	// 	MatchLabels: matchLabels,
	// }

	// selector, _ := toSelector(vvv)

	// clusterLabels := make(map[string]string)
	// clusterLabels["fleet.cattle.io/benchmark"] = "true"

	// vvvvv := selector.Matches(labels.Set(clusterLabels))
	// fmt.Printf("res %v", vvvvv)

	client := rpc_client.New(context.Background())

	res, _ := client.Write("raspberrypi-1", "tm1637", []string{
		"temperature",
		"20",
	})

	fmt.Print(res)

	res1, _ := client.Read("raspberrypi-1", "dht11", []string{})

	fmt.Print(res1)

	client.Close()
}
