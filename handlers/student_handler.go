package handlers

import (
	"go-demo/config"
	"go-demo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateStudent(c *gin.Context) {

	var student models.Student

	if err := c.BindJSON(&student); err != nil {

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	config.DB.Create(&student)

	c.JSON(
		http.StatusCreated,
		student,
	)
}

func GetStudents(c *gin.Context) {
	var students []models.Student
	if err := config.DB.Preload("Teacher").Find(&students).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch students"})
		return
	}
	c.JSON(http.StatusOK, students)
}
