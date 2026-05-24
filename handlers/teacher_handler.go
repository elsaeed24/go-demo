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

func GetTeachers(c *gin.Context) {
	var teachers []models.Teacher
	if err := config.DB.Preload("Students").Find(&teachers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch teachers"})
		return
	}
	c.JSON(http.StatusOK, teachers)
}
