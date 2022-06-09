package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/imhshekhar47/go-rest-api/core"
	"github.com/imhshekhar47/go-rest-api/model"
	"github.com/sirupsen/logrus"
)

type IGreetController interface {
	Hello(ctx *gin.Context)
}

type greetController struct {
	logger *logrus.Entry
}

func newGreetController() *greetController {
	return &greetController{
		logger: core.GetLogger("controller/greetController"),
	}
}

func (ctrl *greetController) Hello(ctx *gin.Context) {
	ctrl.logger.Traceln("entry: Hello()")
	url_path := ctx.Request.URL.Path
	ctx.JSON(http.StatusOK, model.Greet{
		Message: fmt.Sprintf("Hello from %s", url_path),
	})
	ctrl.logger.Traceln("exit: Hello()")
}

// Singleton instance
var greetControllerInstance *greetController

func GetGreetController() *greetController {
	if greetControllerInstance == nil {
		greetControllerInstance = newGreetController()
	}

	return greetControllerInstance
}
