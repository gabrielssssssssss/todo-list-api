package repository

import (
	"fmt"
	"strings"
	"time"

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

	if entity.Description != "" {
		payload = append(payload, fmt.Sprintf("description = '%s'", entity.Description))
	}
	if entity.Status != "" {
		payload = append(payload, fmt.Sprintf("status = '%s'", entity.Status))
	}
	if entity.Title != "" {
		payload = append(payload, fmt.Sprintf("title = '%s'", entity.Title))
	}

	query := fmt.Sprintf(
		`UPDATE tasks SET %s, "updatedAt" = '%s' WHERE id = %s RETURNING "id", "title", "description", "status", "createdAt", "updatedAt";`,
		strings.Join(payload, ", "),
		time.Now().Format("2006-01-02 15:04:05.000000"),
		entity.Id,
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

func (impl *taskRepositoryImpl) GetTasks(entity *entity.TaskPaginationEntity) (*model.TaskPaginationModel, error) {
	var response model.TaskPaginationModel
	var tasks []model.TaskPaginationData

	queryResults := fmt.Sprintf(
		`SELECT id, title, description FROM tasks LIMIT %v OFFSET %v`,
		entity.Limit, entity.Limit*entity.Page,
	)

	rows, err := impl.db.Query(queryResults)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var task model.TaskPaginationData

		err = rows.Scan(&task.Id, &task.Title, &task.Description)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	var count int64
	err = impl.db.QueryRow("SELECT COUNT(*) FROM tasks;").Scan(&count)
	if err != nil {
		return nil, err
	}

	response = model.TaskPaginationModel{
		Data:  tasks,
		Limit: entity.Limit,
		Page:  entity.Page + 1,
		Total: count,
	}

	return &response, nil
}
