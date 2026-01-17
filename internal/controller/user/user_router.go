package user

import (
	"github.com/gin-gonic/gin"
)

func (controller *UserController) Route(rg *gin.RouterGroup) {
	rg.POST("/todos", controller.AddUser)
}
