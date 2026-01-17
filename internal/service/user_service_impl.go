package service

import (
	"fmt"

	"github.com/gabrielssssssssss/todo-list-api/helper"
	"github.com/gabrielssssssssss/todo-list-api/internal/entity"
	"github.com/gabrielssssssssss/todo-list-api/internal/model"
)

func (impl userServiceImpl) AddUser(request *model.UserModel) (*model.UserTokenModel, error) {
	if !helper.IsEmailValid(request.Email) || !helper.IsStrongerPassword(request.Password) {
		return nil, fmt.Errorf("Email or password doesnt match with correct value.")
	}

	hashedPassword, err := helper.HashPassword(request.Password)
	if err != nil {
		return nil, err
	}

	input := entity.UserEntity{
		Name:     request.Name,
		Email:    request.Email,
		Password: hashedPassword,
	}

	response, err := impl.repository.AddUser(&input)
	if err != nil {
		return response, err
	}

	return response, nil
}
