package service

import (
	"fmt"

	"github.com/gabrielssssssssss/todo-list-api/internal/entity"
	"github.com/gabrielssssssssss/todo-list-api/internal/model"
)

func (repository taskServiceImpl) AddTask(request *model.TaskModel) (*model.TaskModel, error) {
	if request.Title == "" || request.Description == "" {
		return nil, fmt.Errorf("Title or description is missing.")
	}

	input := entity.TaskEntity{
		Title:       request.Title,
		Description: request.Description,
	}

	response, err := repository.query.AddTask(&input)
	if err != nil {
		return response, err
	}

	return response, nil
}
