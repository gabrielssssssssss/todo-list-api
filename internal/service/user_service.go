package service

import (
	"github.com/gabrielssssssssss/todo-list-api/internal/model"
	"github.com/gabrielssssssssss/todo-list-api/internal/repository"
)

type UserService interface {
	Register(request *model.UserModel) (*model.UserModelResponse, error)
	Login(request *model.UserModel) (*model.UserModelResponse, error)
}

type userServiceImpl struct {
	repository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userServiceImpl{
		repository: userRepository,
	}
}
