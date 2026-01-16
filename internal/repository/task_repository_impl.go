package repository

import (
	"fmt"

	"github.com/gabrielssssssssss/todo-list-api/config"
	"github.com/gabrielssssssssss/todo-list-api/internal/entity"
	"github.com/gabrielssssssssss/todo-list-api/internal/model"
)

func (impl *taskRepositoryImpl) AddTask(entity *entity.TaskEntity) (*model.TaskModel, error) {
	_, cancel := config.NewPostgresContext()
	defer cancel()

	query := fmt.Sprintf(
		`INSERT INTO tasks (title, description, status) VALUES ('%s', '%s', '%s') RETURNING "id", "title", "description", "status", "createdAt", "updatedAt";`,
		entity.Title, entity.Description, entity.Status,
	)

	var response model.TaskModel
	err := impl.db.QueryRow(query).Scan(
		&response.Id, &response.Title, &response.Description, &response.Status, &response.CreatedAt, &response.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
