package server

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"

	"screencapturer/internal/constant"
	"screencapturer/ui"
)

func NewServer() *http.Server {
	r := registerRoutes()

	return &http.Server{
		Addr:    fmt.Sprintf(":%v", constant.CLIENT_WEB_PORT),
		Handler: r,
	}
}

func registerRoutes() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{
		"http://localhost:5173", // dev ui
		"http://localhost:9999", // prod
	}
	corsConfig.AllowHeaders = []string{
		"Authorization",
		"Content-Type",
		"Access-Control-Allow-Origin",
		"X-CSRF-Token",
	}
	corsConfig.ExposeHeaders = []string{
		"X-Total-Count",
	}
	router.Use(cors.New(corsConfig))

	// UI
	router.Use(static.Serve("/", static.EmbedFolder(ui.EmbeddedFS, "dist")))

	v1 := router.Group("/api")
	{
		computersGroup := v1.Group("/computers")
		registerComputerRoutes(computersGroup)
	}

	router.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/")
	})

	return router
}
