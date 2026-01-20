package task

import (
	"net/http"
	"strconv"

	"github.com/gabrielssssssssss/todo-list-api/internal/entity"
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
	var req entity.TaskEntity

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	req.Token = c.GetHeader("Authorization")

	resp, err := controller.TaskService.AddTask(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (controller *TaskController) DeleteTask(c *gin.Context) {
	input := entity.TaskEntity{
		Id:    c.Param("id"),
		Token: c.GetHeader("Authorization"),
	}

	_, err := controller.TaskService.DeleteTask(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
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
		return
	}

	response, err := controller.TaskService.UpdateTask(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
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
		return
	}

	c.JSON(http.StatusOK, response)
}
