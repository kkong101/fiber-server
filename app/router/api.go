package router

import "github.com/gofiber/fiber/v2"

type Router struct {
	App fiber.Router
}

func NewRouter(fiber *fiber.App) *Router {
	return &Router{
		App: fiber,
	}
}

func (r *Router) Register() {
	r.App.Get("/test", func(ctx *fiber.Ctx) error {
		return ctx.SendString("TEST")
	})

}
