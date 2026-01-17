package repository

import (
	"github.com/gabrielssssssssss/todo-list-api/internal/entity"
	"github.com/gabrielssssssssss/todo-list-api/internal/model"
)

func (impl *userRepositoryImpl) AddUser(entity *entity.UserEntity) (*model.UserModel, error) {
	query := `
		INSERT INTO users (
			name, 
			email,
			password
		) VALUES ($1, $2, $3)`

	var response model.UserModel
	_, err := impl.db.Query(
		query,
		entity.Name,
		entity.Email,
		entity.Password)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
