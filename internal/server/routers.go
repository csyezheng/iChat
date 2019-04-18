package server

import (
	"github.com/csyezheng/iChat/internal/api"
	"github.com/csyezheng/iChat/internal/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func registerRoutes(app *gin.Engine, conf common.Config) {
	// Favicon images
	app.StaticFile("/favicon.ico", conf.HttpFaviconPath()+"/favicon.ico")
	app.StaticFile("/favicon.png", conf.HttpFaviconPath()+"/favicon.png")

	// Static assets like js and css files
	app.Static("/assets", conf.HttpPublicPath())

	// JSON-REST API Version 1
	v1 := app.Group("/api/v1")
	{
		// Handle websocket clients.
		api.ServeWebSocket(v1, conf)
	}

	// Default HTML page
	app.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", conf)
	})
}
