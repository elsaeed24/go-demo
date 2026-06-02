package handlers

import (
	"go-demo/config"
	"go-demo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateStudent(c *gin.Context) {

	var student models.Student

	//if variable := function(); condition {
	//	...
	//}

	//if age := 20; age >= 18 {
	//	fmt.Println("Adult")
	//}
	//nil = كل حاجة تمام

	//err := c.BindJSON(&student)   = var err error err = c.BindJSON(&student)
	//
	//if err != nil {
	//	...
	//}

	if err := c.BindJSON(&student); err != nil {

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	//config.DB.Create(&student)

	result := config.DB.Create(&student)

	if result.Error != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": result.Error.Error(),
			},
		)
		return
	}

	c.JSON(
		http.StatusCreated,
		student,
	)
}

func GetStudents(c *gin.Context) {
	var students []models.Student //create empty slice
	if err := config.DB.Preload("Teacher").Find(&students).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch students"})
		return
	}
	c.JSON(http.StatusOK, students)
}

func UpdateStudent(c *gin.Context) {

	var student models.Student

	id := c.Param("id")

	if err := config.DB.First(&student, id).Error; err != nil {
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"error": "Student not found",
			},
		)
		return
	}

	// استقبل البيانات الجديدة
	if err := c.BindJSON(&student); err != nil {
		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	// حفظ التعديلات
	if err := config.DB.Save(&student).Error; err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		student,
	)
}

func DeleteStudent(c *gin.Context) {

	var student models.Student

	id := c.Param("id")

	if err := config.DB.First(&student, id).Error; err != nil {
		c.JSON(
			http.StatusNotFound,
			gin.H{
				"error": "Student not found",
			},
		)
		return
	}

	if err := config.DB.Delete(&student).Error; err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{
			"message": "Student deleted successfully",
		},
	)
}
