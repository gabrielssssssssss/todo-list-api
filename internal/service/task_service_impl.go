package service

import (
	"fmt"

	"github.com/gabrielssssssssss/todo-list-api/internal/entity"
	"github.com/gabrielssssssssss/todo-list-api/internal/model"
)

func (impl taskServiceImpl) AddTask(request *model.TaskModel) (*model.TaskModelResponse, error) {
	if request.Title == "" || request.Description == "" {
		return nil, fmt.Errorf("Title or description is missing.")
	}

	input := entity.TaskEntity{
		Title:       request.Title,
		Description: request.Description,
	}

	response, err := impl.repository.AddTask(&input)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (impl taskServiceImpl) DeleteTask(request *model.TaskModel) (bool, error) {
	if request.TaskId == "" {
		return false, fmt.Errorf("The value is not number.")
	}

	input := entity.TaskEntity{
		TaskId: request.TaskId,
	}

	response, err := impl.repository.DeleteTask(&input)
	if err != nil {
		return false, err
	}

	return response, nil
}

func (impl taskServiceImpl) UpdateTask(request *model.TaskModel) (*model.TaskModelResponse, error) {
	input := entity.TaskEntity{
		TaskId:      request.TaskId,
		Title:       request.Title,
		Status:      request.Status,
		Description: request.Description,
	}

	response, err := impl.repository.UpdateTask(&input)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (impl taskServiceImpl) GetTasks(request *model.TaskPaginationModel) (*model.TaskPaginationModelResponse, error) {
	input := entity.TaskPaginationEntity{
		Page:  request.Page,
		Limit: request.Limit,
	}

	response, err := impl.repository.GetTasks(&input)
	if err != nil {
		return response, err
	}

	return response, nil
}
