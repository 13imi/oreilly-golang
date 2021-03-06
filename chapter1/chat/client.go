package main

import (
	"github.com/gorilla/websocket"
)

// client is for one person
type client struct {
	// socket is websocket for client
	socket *websocket.Conn
	// send is channel got messages
	send chan []byte
	// room is chat room joining client
	room *room
}

func (c *client) read() {
	for {
		if _, msg, err := c.socket.ReadMessage(); err == nil {
			c.room.forward <- msg
		} else {
			break
		}
	}
	c.socket.Close()
}

func (c *client) write() {
	for msg := range c.send {
		if err := c.socket.WriteMessage(websocket.TextMessage, msg); err != nil {
			break
		}
	}
	c.socket.Close()
}
