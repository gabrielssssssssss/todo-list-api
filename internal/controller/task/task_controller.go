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
	var req model.TaskModel
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "invalid_request",
			"message": err.Error(),
		})
		return
	}

	task, err := controller.TaskService.AddTask(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "task_creation_failed",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "task created successfully",
		"data":    task,
	})
}

func (controller *TaskController) DeleteTask(c *gin.Context) {
	taskID := c.Param("id")
	if taskID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "missing_task_id",
			"message": "task id is required",
		})
		return
	}

	task := model.TaskModel{
		TaskId: taskID,
	}

	if _, err := controller.TaskService.DeleteTask(&task); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "task_deletion_failed",
			"message": err.Error(),
		})
		return
	}

	c.Status(http.StatusNoContent)
}

func (controller *TaskController) UpdateTask(c *gin.Context) {
	taskID := c.Param("id")
	if taskID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "missing_task_id",
			"message": "task id is required",
		})
		return
	}

	var req model.TaskModel
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "invalid_request",
			"message": err.Error(),
		})
		return
	}

	req.TaskId = taskID

	updatedTask, err := controller.TaskService.UpdateTask(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "task_update_failed",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "task updated successfully",
		"data":    updatedTask,
	})
}

func (tc *TaskController) GetTasks(c *gin.Context) {
	limitStr := c.Query("limit")
	pageStr := c.Query("page")

	limit, err := strconv.ParseInt(limitStr, 10, 64)
	if err != nil || limit <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "invalid_limit",
			"message": "limit must be a positive number",
		})
		return
	}

	page, err := strconv.ParseInt(pageStr, 10, 64)
	if err != nil || page <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "invalid_page",
			"message": "page must be a positive number",
		})
		return
	}

	request := model.TaskPaginationModel{
		Limit: limit,
		Page:  page - 1,
	}

	tasks, err := tc.TaskService.GetTasks(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "task_fetch_failed",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "tasks retrieved successfully",
		"data":    tasks,
	})
}
