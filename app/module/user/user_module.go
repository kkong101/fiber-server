package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kkong101/fiber-server/app/module/user/controller"
	"github.com/kkong101/fiber-server/app/module/user/repository"
	"github.com/kkong101/fiber-server/app/module/user/service"
	"go.uber.org/fx"
)

type UserRouter struct {
	App        fiber.Router
	Controller *controller.Controller
}

var NewUserModule = fx.Options(

	fx.Provide(controller.NewUserController),
	fx.Provide(service.NewUserService),
	fx.Provide(repository.NewUserRepository),
)

func (u *UserRouter) RegisterUserRoutes() {
	userController := u.Controller.User

	u.App.Route("/user", func(router fiber.Router) {
		router.Get("/test", userController.GetAllUser)
	})
}
