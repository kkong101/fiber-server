package controller

import "github.com/kkong101/fiber-server/app/module/user/service"

type Controller struct {
	User *UserController
}

func NewController(userService *service.UserService) *Controller {
	return &Controller{User: &UserController{userService}}
}
