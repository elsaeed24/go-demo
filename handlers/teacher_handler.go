package handlers

import (
	"go-demo/config"
	"go-demo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTeacher(c *gin.Context) {

	var teacher models.Teacher

	c.BindJSON(&teacher)

	config.DB.Create(&teacher)

	c.JSON(
		http.StatusCreated,
		teacher,
	)
}
