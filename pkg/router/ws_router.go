package router

import (
	"github.com/darcoprogramador/go-chat/app/controllers"
	"github.com/darcoprogramador/go-chat/app/handlers"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

type WSRouter struct {
	hub controllers.MessageHub
}

func (h WSRouter) InstallRouter(app *fiber.App) {
	app.Get("/ws", websocket.New(handlers.WsClientHandler(h.hub)))
}

func NewWSRouter(hub controllers.MessageHub) *WSRouter {
	return &WSRouter{
		hub: hub,
	}
}
