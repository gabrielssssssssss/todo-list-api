package repository

import (
	"fmt"

	"github.com/gabrielssssssssss/todo-list-api/config"
	"github.com/gabrielssssssssss/todo-list-api/internal/entity"
	"github.com/gabrielssssssssss/todo-list-api/internal/model"
)

func (client *taskImplementation) AddTask(entity *entity.TaskEntity) (*model.TaskModel, error) {
	_, cancel := config.NewPostgresContext()
	defer cancel()

	query := fmt.Sprintf(
		"INSERT INTO tasks(description, status) VALUES (%s, %s)",
		entity.Description, entity.Status,
	)

	response, err := client.db.Query(query)
	if err != nil {
		return nil, err
	}

	fmt.Println(response.Scan("title"))
	return nil, nil
}
