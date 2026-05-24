package routes

import (
	"go-demo/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Server is running",
		})
	})

	router.POST("/students", handlers.CreateStudent)
	router.POST("/teachers", handlers.CreateTeacher)
}
