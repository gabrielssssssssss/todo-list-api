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
	var req model.UserModel
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "invalid_request",
			"message": err.Error(),
		})
		return
	}

	user, err := controller.UserService.Register(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "registration_failed",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "user registered successfully",
		"data":    user,
	})
}

func (controller *UserController) Login(c *gin.Context) {
	var req model.UserModel
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "invalid_request",
			"message": err.Error(),
		})
		return
	}

	token, err := controller.UserService.Login(&req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error":   "login_failed",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "login successful",
		"data":    token,
	})
}
