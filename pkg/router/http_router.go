package router

import (
	"github.com/darcoprogramador/go-chat/app/handlers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type HttpRouter struct {
}

func (h HttpRouter) InstallRouter(app *fiber.App) {
	group := app.Group("", cors.New())
	group.Get("/", handlers.RenderMain)
}

func NewHttpRouter() *HttpRouter {
	return &HttpRouter{}
}
