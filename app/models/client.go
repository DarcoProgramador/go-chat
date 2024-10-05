package models

import (
	"sync"

	"github.com/gofiber/contrib/websocket"
)

// Add more data to this type if needed
type Client struct {
	IsClosing  bool
	Mu         sync.Mutex
	Number     int
	Connection *websocket.Conn
}
