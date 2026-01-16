package task

import (
	"net/http"

	"github.com/gabrielssssssssss/todo-list-api/internal/model"
	"github.com/gabrielssssssssss/todo-list-api/internal/service"
	"github.com/gin-gonic/gin"
)

type TaskController struct {
	TaskService service.TaskService
}

func NewTaskController(TaskService *service.TaskService) TaskController {
	return TaskController{TaskService: *TaskService}
}

func (controller *TaskController) AddTask(c *gin.Context) {
	var request = model.TaskModel{}

	err := c.ShouldBindBodyWithJSON(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}

	response, err := controller.TaskService.AddTask(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, response)
}
