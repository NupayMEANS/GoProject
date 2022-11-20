package main

import (
	"goProject/config"
	"goProject/controller"

	"github.com/gin-gonic/gin"
)

func main() {

	config.Database()

	router := gin.Default()
	router.POST("/createUser", controller.AddUsers)
	router.Run(":8080")

}
