package main

import (
	"go-demo/config"
	"go-demo/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	config.ConnectDB()

	router := gin.Default()

	routes.SetupRoutes(router)

	router.Run(":8080")
}
