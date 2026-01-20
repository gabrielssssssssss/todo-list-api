package service

import (
	"github.com/gabrielssssssssss/todo-list-api/internal/entity"
	"github.com/gabrielssssssssss/todo-list-api/internal/model"
	"github.com/gabrielssssssssss/todo-list-api/internal/repository"
)

type TaskService interface {
	AddTask(request *entity.TaskEntity) (*model.TaskModel, error)
	DeleteTask(request *entity.TaskEntity) (bool, error)
	UpdateTask(request *model.TaskModel) (*model.TaskModel, error)
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
