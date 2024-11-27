package websocket

import (
	"github.com/gorilla/websocket"
)

// Client represents a single chatting user
type Client struct {
	conn *websocket.Conn
	send chan []byte
}