package service

import (
	"fmt"

	"github.com/gabrielssssssssss/todo-list-api/helper"
	"github.com/gabrielssssssssss/todo-list-api/internal/entity"
	"github.com/gabrielssssssssss/todo-list-api/internal/model"
)

func (impl taskServiceImpl) AddTask(request *entity.TaskEntity) (*model.TaskModel, error) {
	if request.Title == "" || request.Description == "" {
		return nil, fmt.Errorf("Title or description is missing.")
	}

	fmt.Println(request.Token)
	ownerId, err := helper.GetJwtValue(request.Token, "owner_id")
	if err != nil {
		return nil, err
	}
	fmt.Println(ownerId)
	input := entity.TaskEntity{
		OwnerId:     ownerId,
		Title:       request.Title,
		Description: request.Description,
	}

	response, err := impl.repository.AddTask(&input)
	if err != nil {
		return response, err
	}

	return response, nil
}

func (impl taskServiceImpl) DeleteTask(request *entity.TaskEntity) (bool, error) {
	if request.Id == "" {
		return false, fmt.Errorf("The value is not number.")
	}

	input := entity.TaskEntity{
		Id: request.Id,
	}

	response, err := impl.repository.DeleteTask(&input)
	if err != nil {
		return false, err
	}

	return response, nil
}

func (impl taskServiceImpl) UpdateTask(request *model.TaskModel) (*model.TaskModel, error) {
	input := entity.TaskEntity{
		Id:          request.Id,
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

func (impl taskServiceImpl) GetTasks(request *model.TaskPaginationModel) (*model.TaskPaginationModel, error) {
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
