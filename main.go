package main

import (
	"loginapi/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/login", controllers.Login)
	router.Run(":8080")
}
