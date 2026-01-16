package controller

import (
	"log"

	"github.com/gabrielssssssssss/todo-list-api/config"
	"github.com/gabrielssssssssss/todo-list-api/internal/controller/task"
	"github.com/gabrielssssssssss/todo-list-api/internal/repository"
	"github.com/gabrielssssssssss/todo-list-api/internal/service"
	"github.com/gin-gonic/gin"
)

func Controller() {
	client, err := config.NewPostgresDatabase()
	if err != nil {
		log.Fatal(err)
	}

	taskRepository := repository.NewTaskRepository(client)
	taskService := service.NewTaskService(taskRepository)
	taskController := task.NewTaskController(&taskService)

	app := gin.Default()
	apiGroup := app.Group("/api")
	taskController.Route(apiGroup)
	app.Run(":8080")
}
