package restlayer

import (
	"fmt"
	"yeelight-server/internal/management"
	"yeelight-server/internal/types"
	"yeelight-server/internal/types/requests"

	"errors"

	"github.com/gin-gonic/gin"
)

var ErrorInvalidRequest = errors.New("Invalid request type")

type LightHandler struct {
	lm *management.YeelightManager
}

func NewLightHandler(lm *management.YeelightManager) *LightHandler {
	return &LightHandler{
		lm: lm,
	}
}

func (lh *LightHandler) RegisterEndpoints(g *gin.Engine) {
	g.PUT("/lights/power", lh.SetPower)
	g.PUT("/lights/brightness", lh.SetBrightness)
	g.PUT("/lights/rgb", lh.SetRgb)
	g.PUT("/lights/hsv", lh.SetHsv)
	g.PUT("/lights/temperature", lh.SetTemperature)
	g.POST("/lights", lh.AddLight)
}

func (lh *LightHandler) GetLights(c *gin.Context) {
	lights := lh.lm.GetLights()
	c.JSON(200, lights)
}

func (lh *LightHandler) AddLight(c *gin.Context) {
	req := new(requests.AddLightRequest)
	if err := c.Bind(&req); err != nil {
		fmt.Println("Bad request ", err)
		c.String(400, "Bad request")
	}
	ids := make([]string, 0)
	for _, addr := range req.LightAddrs {
		light := types.NewYeelight(addr)
		lh.lm.AddLight(light)
		go lh.lm.MonitorLight(addr, light.Id)
		ids = append(ids, light.Id)
	}
	// So this doesn't work. Because it attempts to do this before a TCP connection can be established
	lh.lm.FetchLightStatus(ids)
	c.JSON(200, ids)
}

func (lh *LightHandler) SetPower(c *gin.Context) {
	req := new(requests.SetPowerRequest)
	if err := c.Bind(&req); err != nil {
		fmt.Println("Bad request ", err)
		c.String(400, "Bad request")
	}
	lh.parseRequestAndRunCommand(*req, req.LightIds, c)
}

func (lh *LightHandler) SetBrightness(c *gin.Context) {
	req := new(requests.SetBrightnessRequest)
	if err := c.Bind(&req); err != nil {
		fmt.Println("Bad request ", err)
		c.String(400, "Bad request")
	}
	lh.parseRequestAndRunCommand(*req, req.LightIds, c)
}

func (lh *LightHandler) SetRgb(c *gin.Context) {
	req := new(requests.SetRgbRequest)
	if err := c.Bind(&req); err != nil {
		fmt.Println("Bad request ", err)
		c.String(400, "Bad request")
	}
	lh.parseRequestAndRunCommand(*req, req.LightIds, c)
}

func (lh *LightHandler) SetHsv(c *gin.Context) {
	req := new(requests.SetHsvRequest)
	if err := c.Bind(&req); err != nil {
		fmt.Println("Bad request ", err)
		c.String(400, "Bad request")
	}
	lh.parseRequestAndRunCommand(*req, req.LightIds, c)
}

func (lh *LightHandler) SetTemperature(c *gin.Context) {
	req := new(requests.SetTemperatureRequest)
	if err := c.Bind(&req); err != nil {
		fmt.Println("Bad request ", err)
		c.String(400, "Bad request")
	}
	lh.parseRequestAndRunCommand(*req, req.LightIds, c)
}

func (lh *LightHandler) parseRequestAndRunCommand(req interface{}, lightIds []string, c *gin.Context) {
	cmd, err := createCommand(req)
	if cmd, err = createCommand(req); err != nil {
		c.String(400, "Invalid Command")
		return
	}
	id := lh.lm.RunCommand(cmd, lightIds)
	c.String(200, id)
}

func createCommand(req interface{}) (c *types.Command, err error) {
	switch r := req.(type) {
	default:
		fmt.Printf(ErrorInvalidRequest.Error(), r)
		err = ErrorInvalidRequest
	case requests.SetPowerRequest:
		c = management.NewSetPowerCommand(r.PowerOn, r.Smooth, r.Duration, &r.Mode)
	case requests.SetBrightnessRequest:
		c = management.NewSetBrightnessCommand(r.Brightness, r.Smooth, r.Duration)
	case requests.SetHsvRequest:
		c = management.NewSetHsvCommand(r.Hue, r.Sat, r.Smooth, r.Duration)
	case requests.SetRgbRequest:
		c = management.NewSetRgbCommand(r.Rgb, r.Smooth, r.Duration)
	case requests.SetTemperatureRequest:
		c = management.NewSetCtAbxCommand(r.Ct, r.Smooth, r.Duration)
	}
	return
}
