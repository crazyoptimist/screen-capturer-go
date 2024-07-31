package server

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"screencapturer/internal/constant"
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
	corsConfig.AllowAllOrigins = true
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

	// Health check route
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	v1 := router.Group("/api")
	{
		computersGroup := v1.Group("/computers")
		registerComputerRoutes(computersGroup)
	}

	return router
}
