package restlayer

import (
	"yeelight-server/internal/management"

	"github.com/gin-gonic/gin"
)

type LightManagementRest struct {
	wsHub        *Hub
	lightHandler *LightHandler
}

func NewLightManagementRest(lm *management.YeelightManager, broadcastChannel chan []byte) *LightManagementRest {
	return &LightManagementRest{
		wsHub:        newHub(broadcastChannel),
		lightHandler: NewLightHandler(lm),
	}
}

func (r *LightManagementRest) Run(g *gin.Engine, addr string) {
	go r.wsHub.Run()
	r.lightHandler.RegisterEndpoints(g)
	g.GET("/ws", func(c *gin.Context) {
		serveWs(r.wsHub, c.Writer, c.Request)
	})
	g.Run(addr)
}
