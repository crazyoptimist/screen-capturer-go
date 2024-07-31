package server

import (
	"fmt"
	"net/http"
	"screencapturer/internal/config"

	"github.com/gin-gonic/gin"
)

func NewServer() *http.Server {
	r := registerRoutes()

	return &http.Server{
		Addr:    fmt.Sprintf(":%v", config.CLIENT_WEB_PORT),
		Handler: r,
	}
}

func registerRoutes() *gin.Engine {
	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

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
