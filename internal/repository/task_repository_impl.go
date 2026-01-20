package repository

import (
	"fmt"
	"strings"
	"time"

	"github.com/gabrielssssssssss/todo-list-api/config"
	"github.com/gabrielssssssssss/todo-list-api/internal/entity"
	"github.com/gabrielssssssssss/todo-list-api/internal/model"
)

func (impl *taskRepositoryImpl) AddTask(entity *entity.TaskEntity) (*model.TaskModelResponse, error) {
	_, cancel := config.NewPostgresContext()
	defer cancel()

	query := `
		INSERT INTO tasks (
			title,
			description,
			status,
			owner_id
		)
		VALUES ($1, $2, $3, $4)
		RETURNING
			"id",
			"title",
			"description",
			"status",
			"created_at",
			"updated_at";
	`

	var response model.TaskModelResponse

	err := impl.db.QueryRow(
		query,
		entity.Title,
		entity.Description,
		entity.Status,
		entity.OwnerId,
	).Scan(
		&response.TaskId,
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

	_, err := impl.db.Query(`DELETE FROM tasks WHERE id = $1 AND owner_id = $2;`, entity.TaskId, entity.OwnerId)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (impl *taskRepositoryImpl) UpdateTask(entity *entity.TaskEntity) (*model.TaskModelResponse, error) {
	_, cancel := config.NewPostgresContext()
	defer cancel()

	var (
		payload  []string
		response model.TaskModelResponse
	)

	if entity.Description != "" {
		payload = append(payload, fmt.Sprintf("description = '%s'", entity.Description))
	}
	if entity.Status != "" {
		payload = append(payload, fmt.Sprintf("status = '%s'", entity.Status))
	}
	if entity.Title != "" {
		payload = append(payload, fmt.Sprintf("title = '%s'", entity.Title))
	}

	query := fmt.Sprintf(`
	UPDATE tasks
	SET %s,
	    "updated_at" = $1
	WHERE id = $2
	AND owner_id = $3
	RETURNING
		"id",
		"title",
		"description",
		"status",
		"created_at",
		"updated_at";
	`, strings.Join(payload, ", "))

	err := impl.db.QueryRow(
		query,
		time.Now(),
		entity.TaskId,
		entity.OwnerId,
	).Scan(
		&response.TaskId,
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

func (impl *taskRepositoryImpl) GetTasks(entity *entity.TaskPaginationEntity) (*model.TaskPaginationModelResponse, error) {
	_, cancel := config.NewPostgresContext()
	defer cancel()

	var (
		response model.TaskPaginationModelResponse
		tasks    []model.TaskPaginationModelResults
	)

	query := `
		SELECT
			id,
			title,
			description
		FROM tasks
		WHERE owner_id = $1
		LIMIT $2 OFFSET $3;
	`

	offset := entity.Limit * entity.Page

	rows, err := impl.db.Query(query, entity.OwnerId, entity.Limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var task model.TaskPaginationModelResults

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
		`SELECT COUNT(*) FROM tasks WHERE owner_id = $1;`,
		entity.OwnerId,
	).Scan(&total)
	if err != nil {
		return nil, err
	}

	response = model.TaskPaginationModelResponse{
		Results: tasks,
		Limit:   entity.Limit,
		Page:    entity.Page + 1,
		Total:   total,
	}

	return &response, nil
}
