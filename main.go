package main

import (
	"fmt"
	"net/http"
	"net/url"

	"strings"

	"github.com/gin-gonic/gin"
	"github.com/imhshekhar47/go-rest-api/controller"
	"github.com/imhshekhar47/go-rest-api/core"
)

var (

	// config
	appConfig core.AppConfig = core.GetAppConfig()

	// controller
	actuatorController controller.IActuatorController = controller.GetActuatorController()
	apiGreetController controller.IGreetController    = controller.GetGreetController()
)

var logger = core.GetLogger("main")

func main() {

	logger.Info("entry: main")
	server := gin.New()

	server.GET("/", func(ctx *gin.Context) {
		location := url.URL{
			Path: "/actuator/health",
		}
		ctx.Redirect(http.StatusMovedPermanently, location.RequestURI())
	})

	actuatorRoutes := server.Group("/actuator")
	{
		actuatorRoutes.GET("/health", actuatorController.Health)
		actuatorRoutes.GET("/info", actuatorController.Info)
	}

	apiRoutes := server.Group(fmt.Sprintf("%s/api", appConfig.Server.BasePath))
	{
		apiRoutes.GET("/greet", apiGreetController.Hello)
	}

	logger.Infof("Starting server in %s mode on 0.0.0.0:%s", appConfig.Server.Mode, appConfig.Server.Port)

	if strings.ToUpper(appConfig.Server.Mode) == "RELEASE" {
		gin.SetMode(gin.ReleaseMode)
	}

	server.Run(fmt.Sprintf("0.0.0.0:%s", appConfig.Server.Port))
}
