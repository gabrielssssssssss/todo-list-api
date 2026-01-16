package repository

import (
	"fmt"
	"reflect"
	"strings"

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

func (impl *taskRepositoryImpl) DeleteTask(entity *entity.TaskEntity) (bool, error) {
	_, cancel := config.NewPostgresContext()
	defer cancel()

	query := fmt.Sprintf(
		`DELETE FROM tasks WHERE id = %s;`,
		entity.Id,
	)

	_, err := impl.db.Query(query)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (impl *taskRepositoryImpl) UpdateTask(entity *entity.TaskEntity) (*model.TaskModel, error) {
	_, cancel := config.NewPostgresContext()
	defer cancel()

	var payload []string

	k := reflect.TypeOf(entity)
	v := reflect.ValueOf(entity)

	for i := 0; i < k.NumField(); i++ {
		key, value := v.Field(i).Interface(), k.Field(i).Name
		if value != "" {
			payload = append(payload, fmt.Sprintf("%s = %s", key, value))
		}
	}

	query := fmt.Sprintf(
		`UPDATE tasks SET %s WHERE id = %s;`,
		strings.Join(payload, ","), entity.Id,
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
