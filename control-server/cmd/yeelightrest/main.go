package main

import (
	"fmt"
	"yeelight-control-server/internal/discovery"
)

func main() {
	discoverService := discovery.DiscoveryService{}
	c := discoverService.Start()
	for light := range c {
		fmt.Println(light)
	}
}
