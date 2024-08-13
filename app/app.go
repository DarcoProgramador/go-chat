package app

import (
	"github.com/darcoprogramador/go-chat/app/controllers"
	"github.com/darcoprogramador/go-chat/pkg/router"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
)

func NewApplication(hub controllers.MessageHub) *fiber.App {
	// HTML template engine
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{Views: engine})

	// Middlewares
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) { // Returns true if the client requested upgrade to the WebSocket protocol
			return c.Next()
		}
		return c.SendStatus(fiber.StatusUpgradeRequired)
	})
	// routes
	router.InstallRouter(app, hub)
	return app
}
