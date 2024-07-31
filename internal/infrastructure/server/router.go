package server

import (
	"github.com/gin-gonic/gin"

	"screencapturer/internal/config"
	"screencapturer/internal/infrastructure/controller"
)

func registerComputerRoutes(g *gin.RouterGroup) {
	controllers := controller.NewComputerController(config.DB)
	g.GET("", controllers.FindAll)
	g.GET(":id", controllers.FindById)
	g.POST("", controllers.Create)
	g.PATCH(":id", controllers.Update)
	g.DELETE(":id", controllers.Delete)
}
