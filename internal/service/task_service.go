package service

import (
	"github.com/gabrielssssssssss/todo-list-api/internal/entity"
	"github.com/gabrielssssssssss/todo-list-api/internal/model"
	"github.com/gabrielssssssssss/todo-list-api/internal/repository"
)

type TaskService interface {
	AddTask(request *model.TaskModel) (*entity.TaskEntity, error)
	DeleteTask(request *model.TaskModel) (bool, error)
	UpdateTask(request *model.TaskModel) (*entity.TaskEntity, error)
	GetTasks(request *model.TaskPaginationModel) (*model.TaskPaginationModel, error)
}

type taskServiceImpl struct {
	repository repository.TaskRepository
}

func NewTaskService(taskRepository repository.TaskRepository) TaskService {
	return &taskServiceImpl{
		repository: taskRepository,
	}
}
