package main

import (
	"yeelight-server/internal/discovery"
	"yeelight-server/internal/management"
	"yeelight-server/internal/restlayer"

	"github.com/gin-gonic/gin"
)

func main() {
	discoverService := discovery.NewDiscoveryService()
	lightManager := management.NewYeelightManager()
	go lightManager.Start(discoverService.Start())

	managementRest := restlayer.NewLightManagementRest(lightManager)
	route := gin.Default()

	managementRest.Run(route, ":8000")
}
