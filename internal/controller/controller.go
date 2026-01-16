package controller

import "github.com/gin-gonic/gin"

func Controller() {
	app := gin.Default()
	app.Run(":8080")
}
