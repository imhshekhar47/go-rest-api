package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/imhshekhar47/go-rest-api/core"
	"github.com/imhshekhar47/go-rest-api/model"
)

var logger = core.GetLogger("controller")

// Interface definition
type IActuatorController interface {
	Health(ctx *gin.Context)
	Info(ctx *gin.Context)
}

// Interface Implementation
type actuatorController struct{}

func (ctrl *actuatorController) Health(ctx *gin.Context) {
	logger.Tracef("entry: Health")
	ctx.JSON(200, model.Health{
		Status: "UP",
	})
	logger.Tracef("exit: Health")
}

func (ctrl *actuatorController) Info(ctx *gin.Context) {
	logger.Tracef("entry: Info")
	ctx.JSON(200, model.Info{
		Name:    core.GetAppConfig().Application.Name,
		Version: core.GetAppConfig().Application.Version,
	})
	logger.Tracef("exit: Info")
}

// Singleton impl
var instance *actuatorController

func GetActuatorController() IActuatorController {
	if nil == instance {
		instance = new(actuatorController)
	}
	return instance
}
