package controller

import "github.com/kkong101/fiber-server/app/module/user/service"

type UserController struct {
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func getAllUser() error {

}
