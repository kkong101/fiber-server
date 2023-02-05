package service

import "github.com/kkong101/fiber-server/app/module/user/repository"

type UserService struct {
	Repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		Repo: repo,
	}
}
