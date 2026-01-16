package service

import (
	"github.com/gabrielssssssssss/todo-list-api/internal/model"
	"github.com/gabrielssssssssss/todo-list-api/internal/repository"
)

type TaskService interface {
	AddTask(request *model.TaskModel) (*model.TaskModel, error)
}

type taskServiceImpl struct {
	repository repository.TaskRepository
}

func NewTaskService(taskRepository repository.TaskRepository) TaskService {
	return &taskServiceImpl{
		repository: taskRepository,
	}
}
