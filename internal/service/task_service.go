package service

import (
	"github.com/gabrielssssssssss/todo-list-api/internal/model"
	"github.com/gabrielssssssssss/todo-list-api/internal/repository"
)

type TaskService interface {
	AddTask(request *model.TaskModel) (*model.TaskModel, error)
}

type taskServiceImpl struct {
	query repository.TaskRepository
}

func NewService(taskRepository repository.TaskRepository) TaskService {
	return &taskServiceImpl{
		query: taskRepository,
	}
}
