package repository

import (
	"database/sql"

	"github.com/gabrielssssssssss/todo-list-api/internal/entity"
	"github.com/gabrielssssssssss/todo-list-api/internal/model"
)

type TaskRepository interface {
	AddTask(entity *entity.TaskEntity) (*model.TaskModelResponse, error)
	DeleteTask(entity *entity.TaskEntity) (bool, error)
	UpdateTask(entity *entity.TaskEntity) (*model.TaskModelResponse, error)
	GetTasks(entity *entity.TaskPaginationEntity) (*model.TaskPaginationModelResponse, error)
}

type taskRepositoryImpl struct {
	db *sql.DB
}

func NewTaskRepository(client *sql.DB) TaskRepository {
	return &taskRepositoryImpl{
		db: client,
	}
}
