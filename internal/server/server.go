package server

import (
	"fmt"
	"github.com/csyezheng/iChat/internal/common"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

const (
	// idleSessionTimeout defines duration of being idle before terminating a session.
	idleSessionTimeout = time.Second * 55
)

// Start the REST API server using the configuration provided
func Start(conf common.Config) {
	if conf.HttpServerMode() != "" {
		gin.SetMode(conf.HttpServerMode())
	} else if conf.Debug() == false {
		gin.SetMode(gin.ReleaseMode)
	}

	// The hub (the main message router)
	hub := conf.NewHub()
	go hub.Run()

	app := gin.Default()

	// Set template directory
	app.LoadHTMLGlob(conf.HttpTemplatesPath() + "/*")

	registerRoutes(app, conf)
	log.Println("----")
	log.Println(fmt.Sprintf("%s:%d", conf.HttpServerHost(), conf.HttpServerPort()))
	log.Println("----")
	err := app.Run(fmt.Sprintf("%s:%d", conf.HttpServerHost(), conf.HttpServerPort()))
	if err != nil {
		panic(err)
	}
}
