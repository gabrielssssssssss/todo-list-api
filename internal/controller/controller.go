package controller

import (
	"log"

	"github.com/gabrielssssssssss/todo-list-api/config"
	"github.com/gabrielssssssssss/todo-list-api/internal/controller/task"
	"github.com/gabrielssssssssss/todo-list-api/internal/controller/user"
	"github.com/gabrielssssssssss/todo-list-api/internal/middlewares"
	"github.com/gabrielssssssssss/todo-list-api/internal/repository"
	"github.com/gabrielssssssssss/todo-list-api/internal/service"
	"github.com/gin-gonic/gin"
)

func Controller() {
	client, err := config.NewPostgresDatabase()
	if err != nil {
		log.Fatal(err)
	}

	userRepository := repository.NewUserRepository(client)
	taskRepository := repository.NewTaskRepository(client)

	userService := service.NewUserService(userRepository)
	taskService := service.NewTaskService(taskRepository)

	userController := user.NewUserController(&userService)
	taskController := task.NewTaskController(&taskService)

	app := gin.Default()
	app.Use(middlewares.CORSMiddleware())

	apiGroup := app.Group("/api")
	userController.Route(apiGroup)
	taskController.Route(apiGroup)
	app.Run(":8080")
}
