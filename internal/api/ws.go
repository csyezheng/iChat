package api

import (
	"github.com/csyezheng/iChat/internal/common"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// Allow connections from any Origin
	CheckOrigin: func(r *http.Request) bool { return true },
}

// Handles websocket requests from peers
func ServeWebSocket(router *gin.RouterGroup, conf common.Config) {
	router.GET("/ws", func(c *gin.Context) {

		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if _, ok := err.(websocket.HandshakeError); ok {
			log.Println("ws: Not a websocket handshake")
			return
		} else if err != nil {
			log.Println("ws: failed to Upgrade ", err)
			return
		}

		// Register our new client
		cli := &common.Client{Hub: conf.GetHub(), WSConn: conn, Send: make(chan common.Message)}
		cli.GetHub().Register() <- cli

		// Make sure we close the connection when the function returns
		// Use conn to send and receive messages.
		// Do work in goroutines to return from serveWebSocket() to release file pointers.
		// Otherwise "too many open files" will happen. ie. collect memory in new goroutines
		go cli.WritePump()
		go cli.ReadPump()
	})
}
