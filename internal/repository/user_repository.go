package repository

import (
	"database/sql"

	"github.com/gabrielssssssssss/todo-list-api/internal/entity"
	"github.com/gabrielssssssssss/todo-list-api/internal/model"
)

type UserRepository interface {
	Register(*entity.UserEntity) (*model.UserModelResponse, error)
	Login(*entity.UserEntity) (*model.UserModelResponse, error)
}

type userRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(client *sql.DB) UserRepository {
	return &userRepositoryImpl{
		db: client,
	}
}
