package task

import "github.com/gin-gonic/gin"

func (controller *TaskController) Route(rg *gin.RouterGroup) {
	rg.POST("/todos", controller.AddTask)
}
