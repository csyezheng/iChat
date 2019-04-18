package common

import (
	"github.com/gorilla/websocket"
	"log"
	"time"
)


const (
	NONE = iota
	WEBSOCK
	LPOLL
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

// Client is a middleman between a websocket connection and the hub or a long polling and the hub.
// A user may have multiple clients.
type Client struct {
	Hub *Hub

	// The websocket connection. Set only for websocket clients
	WSConn *websocket.Conn

	// Buffered channel of outbound messages.
	// The content must be serialized in format suitable for the session.
	Send chan Message
}


// readPump pumps messages from the websocket connection to the hub.
//
// The application runs readPump in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.
func (c *Client) ReadPump() {
	defer func() {
		c.Hub.unregister <- c
		c.WSConn.Close()
	}()
	c.WSConn.SetReadLimit(maxMessageSize)
	c.WSConn.SetReadDeadline(time.Now().Add(pongWait))
	c.WSConn.SetPongHandler(func(string) error { c.WSConn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		var message Message
		err := c.WSConn.ReadJSON(&message)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		c.Hub.broadcast <- message
	}
}

// writePump pumps messages from the hub to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *Client) WritePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.WSConn.Close()
	}()
	for {
		select {
		case message, ok := <-c.Send:
			c.WSConn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.WSConn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			err := c.WSConn.WriteJSON(message)
			if err != nil {
				return
			}

		case <-ticker.C:
			c.WSConn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.WSConn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (c *Client) GetHub() *Hub {
	return c.Hub
}