package restlayer

import (
	"yeelight-control-server/internal/management"
	"yeelight-control-server/internal/types"

	"github.com/gin-gonic/gin"
)

type LightHandler struct {
	lm *management.YeelightManager
}

func NewLightHandler(lm *management.YeelightManager) *LightHandler {
	return &LightHandler{
		lm: lm,
	}
}

func (lm *LightHandler) TurnOffLight(c *gin.Context) {
	lights := make([]string, 0)
	if err := c.BindQuery(&lights); err != nil {
		c.String(400, "Missing lights query param")
	}
	cmd := management.NewSetPowerCommand(false, true, 1000, &types.DefaultLightMode)
	id := lm.lm.RunCommand(cmd, lights)
	c.String(200, id)
}
