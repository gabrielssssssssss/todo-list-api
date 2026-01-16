package main

import (
	"github.com/gabrielssssssssss/todo-list-api/internal/controller"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	controller.Controller()
}
