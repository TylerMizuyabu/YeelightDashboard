package main

import (
	"yeelight-server/internal/discovery"
	"yeelight-server/internal/management"
	"yeelight-server/internal/restlayer"

	"github.com/gin-gonic/gin"
)

func main() {
	broadcastChannel := make(chan []byte)
	discoverService := discovery.NewDiscoveryService()
	lightManager := management.NewYeelightManager()
	go lightManager.Start(discoverService.Start(), broadcastChannel)

	managementRest := restlayer.NewLightManagementRest(lightManager, broadcastChannel)
	route := gin.Default()

	managementRest.Run(route, ":8000")
}
