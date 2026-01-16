package repository

import (
	"database/sql"

	"github.com/gabrielssssssssss/todo-list-api/internal/entity"
	"github.com/gabrielssssssssss/todo-list-api/internal/model"
)

type TaskRepository interface {
	AddTask(entity *entity.TaskEntity) (*model.TaskModel, error)
}

type taskRepositoryImpl struct {
	db *sql.DB
}

func NewTaskRepository(client *sql.DB) TaskRepository {
	return &taskRepositoryImpl{
		db: client,
	}
}
