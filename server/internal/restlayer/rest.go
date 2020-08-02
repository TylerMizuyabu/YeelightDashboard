package restlayer

import (
	"yeelight-server/internal/management"

	"github.com/gin-gonic/gin"
)

type LightManagementRest struct {
	lightHandler *LightHandler
}

func NewLightManagementRest(lm *management.YeelightManager) *LightManagementRest {
	return &LightManagementRest{
		lightHandler: NewLightHandler(lm),
	}
}

func (r *LightManagementRest) Run(g *gin.Engine, addr string) {
	r.lightHandler.RegisterEndpoints(g)
	g.Run(addr)
}
