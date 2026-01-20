package user

import (
	"net/http"

	"github.com/gabrielssssssssss/todo-list-api/internal/model"
	"github.com/gabrielssssssssss/todo-list-api/internal/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(UserService *service.UserService) UserController {
	return UserController{UserService: *UserService}
}

func (controller *UserController) Register(c *gin.Context) {
	var request = model.UserModel{}

	err := c.ShouldBindBodyWithJSON(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}

	response, err := controller.UserService.Register(&request)
	if err == nil {
		c.JSON(http.StatusOK, response)
	}

	c.JSON(http.StatusInternalServerError, gin.H{
		"message": err.Error(),
	})
}

func (controller *UserController) Login(c *gin.Context) {
	var request = model.UserModel{}

	err := c.ShouldBindBodyWithJSON(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}

	response, err := controller.UserService.Login(&request)
	if err == nil {
		c.JSON(http.StatusOK, response)
	}

	c.JSON(http.StatusInternalServerError, gin.H{
		"message": err.Error(),
	})
}
