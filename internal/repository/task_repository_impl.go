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

	query := `
		INSERT INTO tasks (
			title,
			description,
			status
		)
		VALUES ($1, $2, $3)
		RETURNING
			"id",
			"title",
			"description",
			"status",
			"createdAt",
			"updatedAt";
	`

	var response model.TaskModel

	err := impl.db.QueryRow(
		query,
		entity.Title,
		entity.Description,
		entity.Status,
	).Scan(
		&response.Id,
		&response.Title,
		&response.Description,
		&response.Status,
		&response.CreatedAt,
		&response.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (impl *taskRepositoryImpl) DeleteTask(entity *entity.TaskEntity) (bool, error) {
	_, cancel := config.NewPostgresContext()
	defer cancel()

	_, err := impl.db.Query(`DELETE FROM tasks WHERE id = ?;`, entity.Id)
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

	var response model.TaskModel

	query := `
	UPDATE tasks
	SET %s,
	    "updatedAt" = $2
	WHERE id = $3
	RETURNING
		"id",
		"title",
		"description",
		"status",
		"createdAt",
		"updatedAt";
	`

	err := impl.db.QueryRow(
		fmt.Sprintf(query, strings.Join(payload, ", ")),
		time.Now(),
		entity.Id,
	).Scan(
		&response.Id,
		&response.Title,
		&response.Description,
		&response.Status,
		&response.CreatedAt,
		&response.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &response, nil
}

func (impl *taskRepositoryImpl) GetTasks(entity *entity.TaskPaginationEntity) (*model.TaskPaginationModel, error) {
	var (
		response model.TaskPaginationModel
		tasks    []model.TaskPaginationData
	)

	query := `
		SELECT
			id,
			title,
			description
		FROM tasks
		LIMIT $1 OFFSET $2;
	`

	offset := entity.Limit * entity.Page

	rows, err := impl.db.Query(query, entity.Limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var task model.TaskPaginationData

		if err := rows.Scan(
			&task.Id,
			&task.Title,
			&task.Description,
		); err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	var total int64
	err = impl.db.QueryRow(
		`SELECT COUNT(*) FROM tasks;`,
	).Scan(&total)
	if err != nil {
		return nil, err
	}

	response = model.TaskPaginationModel{
		Data:  tasks,
		Limit: entity.Limit,
		Page:  entity.Page + 1,
		Total: total,
	}

	return &response, nil
}
