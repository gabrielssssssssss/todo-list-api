package task

import (
	"github.com/gabrielssssssssss/todo-list-api/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func (controller *TaskController) Route(rg *gin.RouterGroup) {
	rg.Use(middlewares.VerifyJwtToken)
	rg.POST("/todos", controller.AddTask)
	rg.DELETE("/todos/:id", controller.DeleteTask)
	rg.PUT("/todos/:id", controller.UpdateTask)
	rg.GET("/todos", controller.GetTasks)
}
