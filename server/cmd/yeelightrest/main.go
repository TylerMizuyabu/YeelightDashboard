package main

import (
	"yeelight-control-server/internal/discovery"
	"yeelight-control-server/internal/management"
	"yeelight-control-server/internal/restlayer"

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
