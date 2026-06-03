package handlers

import (
	"go-demo/config"
	"go-demo/dto"
	"go-demo/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateStudent(c *gin.Context) {

	var input dto.CreateStudentInput // بنستخدم DTO مش model

	// Bind + Validation
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(), // بيرجع validation errors
		})
		return
	}

	var teacher models.Teacher

	if err := config.DB.First(&teacher, input.TeacherID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Teacher not found",
		})
		return
	}

	// نحول من DTO → Model
	student := models.Student{
		Name:      input.Name,
		Age:       input.Age,
		TeacherID: input.TeacherID,
	}

	// نحفظ في DB
	if err := config.DB.Create(&student).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Response
	c.JSON(http.StatusCreated, gin.H{
		"message": "Student created",
		"data":    student,
	})
}

func GetStudents(c *gin.Context) {

	var students []models.Student // slice فاضية

	// Preload = eager loading للعلاقة
	if err := config.DB.Preload("Teacher").Find(&students).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch students",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": students,
	})
}

func UpdateStudent(c *gin.Context) {

	var student models.Student

	id := c.Param("id")

	// نجيب الطالب من DB
	if err := config.DB.First(&student, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Student not found",
		})
		return
	}

	var input dto.UpdateStudentInput

	// Bind request
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// تحديث القيم (بشكل آمن)
	if input.Name != "" {
		student.Name = input.Name
	}

	if input.Age != 0 {
		student.Age = input.Age
	}

	if input.TeacherID != 0 {
		student.TeacherID = input.TeacherID
	}

	// حفظ
	if err := config.DB.Save(&student).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Student updated",
		"data":    student,
	})
}

func DeleteStudent(c *gin.Context) {

	var student models.Student

	id := c.Param("id")

	// نتاكد انه موجود
	if err := config.DB.First(&student, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Student not found",
		})
		return
	}

	// حذف
	if err := config.DB.Delete(&student).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Student deleted successfully",
	})
}
