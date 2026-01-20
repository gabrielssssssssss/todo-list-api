package repository

import (
	"fmt"
	"os"

	"github.com/gabrielssssssssss/todo-list-api/helper"
	"github.com/gabrielssssssssss/todo-list-api/internal/entity"
	"github.com/gabrielssssssssss/todo-list-api/internal/model"
)

var (
	JWT_SECRET_KEY = []byte(os.Getenv("JWT_SECRET_KEY"))
)

func (impl *userRepositoryImpl) Register(entity *entity.UserEntity) (*model.UserTokenModel, error) {
	query := `
		INSERT INTO users (
			name, 
			email,
			password
		) VALUES ($1, $2, $3)
		RETURNING
			"id"`

	err := impl.db.QueryRow(
		query,
		entity.Name,
		entity.Email,
		entity.Password).Scan(&entity.Id)
	if err != nil {
		return nil, err
	}

	token, err := helper.GenerateJwtToken(entity.Id, entity.Email, JWT_SECRET_KEY)
	if err != nil {
		return nil, err
	}

	return &model.UserTokenModel{Token: token}, nil
}

func (impl *userRepositoryImpl) Login(entity *entity.UserEntity) (*model.UserTokenModel, error) {
	var hashPassword *string
	err := impl.db.QueryRow(
		`SELECT id, password FROM users WHERE email = $1`,
		entity.Email,
	).Scan(&entity.Id, &hashPassword)
	if err != nil {
		return nil, err
	}

	if !helper.VerifyPassword(entity.Password, *hashPassword) {
		return nil, fmt.Errorf("Invalid credentials.")
	}

	token, err := helper.GenerateJwtToken(entity.Id, entity.Email, JWT_SECRET_KEY)
	if err != nil {
		return nil, err
	}

	return &model.UserTokenModel{Token: token}, nil
}
