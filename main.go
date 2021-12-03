package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/imhshekhar47/go-rest-api/controller"
	"github.com/imhshekhar47/go-rest-api/core"
	"log"
	"strings"
)

var (
	appConfig core.AppConfig = core.GetAppConfig()
	actuatorController controller.IActuatorController = controller.GetActuatorController()
)

func main() {
	log.Println("entry: main")
	server := gin.New()

	actuatorRoutes := server.Group("/actuator")
	{
		actuatorRoutes.GET("/health", actuatorController.Health)
		actuatorRoutes.GET("/info", actuatorController.Info)
	}

	log.Printf(fmt.Sprintf("Starting server in %s mode", appConfig.Server.Mode))
	if "RELEASE" == strings.ToUpper(appConfig.Server.Mode) {
		gin.SetMode(gin.ReleaseMode)
	}
	server.Run(fmt.Sprintf("0.0.0.0:%s", appConfig.Server.Port))
}
