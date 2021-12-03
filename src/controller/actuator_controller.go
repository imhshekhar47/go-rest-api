package controller

import(
	"github.com/gin-gonic/gin"
	"github.com/imhshekhar47/go-rest-api/src/core"
	"github.com/imhshekhar47/go-rest-api/src/model"
)
// Interface definition
type IActuatorController interface {
	Health(ctx *gin.Context)
	Info(ctx *gin.Context)
}

// Interface Implementation
type actuatorController struct {}

func (ctrl *actuatorController) Health(ctx *gin.Context) {
	ctx.JSON(200, model.Health{
		Status: "UP",
	})
}

func (ctrl *actuatorController) Info(ctx *gin.Context) {
	ctx.JSON(200, model.Info{
		Name: core.GetAppConfig().Application.Name,
		Version: core.GetAppConfig().Application.Version,
	})
}

// Singleton impl
var instance *actuatorController

func GetActuatorController() IActuatorController {
	if nil == instance {
		instance = new(actuatorController)
	}
	return instance
}
