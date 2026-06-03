package services

import (
	"errors"

	"go-demo/config"
	"go-demo/dto"
	"go-demo/models"

	"gorm.io/gorm"
)

type StudentService struct{}

// NewStudentService = service := StudentService{}
// ليه عملناه أصلاً؟
// عشان نجمع كل العمليات الخاصة بالطلاب:
// Create()
// Update()
// Delete()
// GetAll()

// NewStudentService Constructor
func NewStudentService() *StudentService {
	return &StudentService{}
}

func (s *StudentService) Create(

	input dto.CreateStudentInput,

) (*models.Student, error) {

	var teacher models.Teacher

	if err := config.DB.First(&teacher, input.TeacherID).Error; err != nil {

		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("teacher not found")
		}

		return nil, err
	}

	student := models.Student{
		Name:      input.Name,
		Age:       input.Age,
		TeacherID: input.TeacherID,
	}

	if err := config.DB.Create(&student).Error; err != nil {
		return nil, err
	}

	return &student, nil
}
