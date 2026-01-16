package main

import (
	"fmt"
	"log"

	"github.com/gabrielssssssssss/todo-list-api/config"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	client, err := config.NewPostgresDatabase()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(client.Stats())
}
