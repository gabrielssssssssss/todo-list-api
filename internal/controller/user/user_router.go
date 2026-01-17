package user

import (
	"github.com/gin-gonic/gin"
)

func (controller *UserController) Route(rg *gin.RouterGroup) {
	rg.POST("/register", controller.Register)
	rg.POST("/login", controller.Login)
}
