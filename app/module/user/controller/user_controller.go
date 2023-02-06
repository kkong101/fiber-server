package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kkong101/fiber-server/app/module/user/service"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (con *UserController) GetAllUser(c *fiber.Ctx) error {
	return nil
}
