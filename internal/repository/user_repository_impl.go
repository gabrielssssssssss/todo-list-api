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
		) VALUES (
			$1, 
			$2, 
			$3
		)
		RETURNING
			"id"
			"name"
			"email"
			"password"
			"createdAt"
			"updatedAt";
	`

	var response model.UserModel
	err := impl.db.QueryRow(
		query,
		entity.Name,
		entity.Email,
		entity.Password).Scan(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
