package task

import (
	"net/http"
	"strconv"

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
		c.Abort()
	}

	response, err := controller.TaskService.AddTask(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		c.Abort()
	}

	c.JSON(http.StatusOK, response)
}

func (controller *TaskController) DeleteTask(c *gin.Context) {
	var id = c.Param("id")

	input := model.TaskModel{
		Id: id,
	}

	_, err := controller.TaskService.DeleteTask(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		c.Abort()
	}

	c.Status(204)
}

func (controller *TaskController) UpdateTask(c *gin.Context) {
	var request model.TaskModel

	request.Id = c.Param("id")
	err := c.ShouldBindBodyWithJSON(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		c.Abort()
	}

	response, err := controller.TaskService.UpdateTask(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		c.Abort()
	}

	c.JSON(http.StatusOK, response)
}

func (controller *TaskController) GetTasks(c *gin.Context) {
	var request model.TaskPaginationModel

	limit, err := strconv.ParseInt(c.Query("limit"), 10, 64)
	page, err := strconv.ParseInt(c.Query("page"), 10, 64)

	request.Limit = limit
	request.Page = (page - 1)

	response, err := controller.TaskService.GetTasks(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		c.Abort()
	}

	c.JSON(http.StatusOK, response)
}
