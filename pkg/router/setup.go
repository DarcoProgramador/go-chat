package router

import (
	"github.com/darcoprogramador/go-chat/app/controllers"
	"github.com/gofiber/fiber/v2"
)

func InstallRouter(app *fiber.App, hub controllers.MessageHub) {
	setup(app, NewHttpRouter(), NewWSRouter(hub))
}
func setup(app *fiber.App, router ...Router) {
	for _, r := range router {
		r.InstallRouter(app)
	}
}
