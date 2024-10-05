package handlers

import (
	"log"

	"github.com/darcoprogramador/go-chat/app/controllers"
	"github.com/gofiber/contrib/websocket"
)

func WsClientHandler(hub controllers.MessageHub) func(c *websocket.Conn) {
	return func(c *websocket.Conn) {
		// When the function returns, unregister the client and close the connection
		defer hub.Unregister(c)

		// Register the client/
		hub.Register(c)

		for {
			messageType, message, err := c.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					log.Println("read error:", err)
				}

				return // Calls the deferred function, i.e. closes the connection on error
			}

			if messageType == websocket.TextMessage {
				// Broadcast the received message
				hub.SendMessage(c, string(message))
			} else {
				log.Println("websocket message received of type", messageType)
			}
		}
	}
}
