package controllers

import (
	"encoding/json"
	"log"

	"github.com/darcoprogramador/go-chat/app/models"
	"github.com/gofiber/contrib/websocket"
)

type MessageHub interface {
	RunHub()
	Register(*websocket.Conn)
	Unregister(*websocket.Conn)
	Broadcast(string)
	SendMessage(*websocket.Conn, string)
}

type Hub struct {
	clients    map[*websocket.Conn]*models.Client
	register   chan *websocket.Conn
	broadcast  chan string
	unregister chan *websocket.Conn
}

func NewHub() MessageHub {
	return &Hub{
		clients:    make(map[*websocket.Conn]*models.Client),
		register:   make(chan *websocket.Conn),
		broadcast:  make(chan string),
		unregister: make(chan *websocket.Conn),
	}
}

func (h *Hub) RunHub() {
	count := 1
	for {
		select {
		case connection := <-h.register:
			h.clients[connection] = &models.Client{Number: count}
			count++
			log.Println("connection registered")

		case message := <-h.broadcast:
			log.Println("message received:", message)
			// Send the message to all clients
			for connection, c := range h.clients {
				go func(connection *websocket.Conn, c *models.Client) { // send to each client in parallel so we don't block on a slow client
					c.Mu.Lock()
					defer c.Mu.Unlock()
					if c.IsClosing {
						return
					}
					if err := connection.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
						c.IsClosing = true
						log.Println("write error:", err)

						connection.WriteMessage(websocket.CloseMessage, []byte{})
						connection.Close()
						h.unregister <- connection
					}
				}(connection, c)
			}

		case connection := <-h.unregister:
			// Remove the client from the hub
			delete(h.clients, connection)

			log.Println("connection unregistered")
		}
	}
}

func (h *Hub) Register(connection *websocket.Conn) {
	h.register <- connection
}

func (h *Hub) Unregister(connection *websocket.Conn) {
	h.unregister <- connection
	connection.Close()
}

func (h *Hub) Broadcast(message string) {
	h.broadcast <- message
}

func (h *Hub) SendMessage(connection *websocket.Conn, message string) {
	user := h.clients[connection].Number

	jsonMessage, err := json.Marshal(models.Message{User: user, Message: message})

	if err != nil {
		log.Println("error marshalling message:", err)
		return
	}

	h.Broadcast(string(jsonMessage))
}
