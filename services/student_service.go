package services

import (
	"errors"

	"go-demo/config"
	"go-demo/dto"
	"go-demo/models"

	"gorm.io/gorm"
)

// StudentService مسؤول عن كل العمليات الخاصة بالطلاب
type StudentService struct{}

// NewStudentService Constructor لإنشاء instance من الـ service
func NewStudentService() *StudentService {
	return &StudentService{}
}

// Create Student
func (s *StudentService) Create(
	input dto.CreateStudentInput,
) (*models.Student, error) {

	// متغير لتخزين بيانات المدرس
	var teacher models.Teacher

	// التأكد أن المدرس موجود
	if err := config.DB.First(&teacher, input.TeacherID).Error; err != nil {

		// لو المدرس غير موجود
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("teacher not found")
		}

		// أي خطأ آخر من قاعدة البيانات
		return nil, err
	}

	// تحويل DTO إلى Model
	student := models.Student{
		Name:      input.Name,
		Age:       input.Age,
		TeacherID: input.TeacherID,
	}

	// حفظ الطالب في قاعدة البيانات
	if err := config.DB.Create(&student).Error; err != nil {
		return nil, err
	}

	// إرجاع الطالب بعد الحفظ
	return &student, nil
}

// Update Student
func (s *StudentService) Update(
	id string,
	input dto.UpdateStudentInput,
) (*models.Student, error) {

	// متغير لتخزين الطالب
	var student models.Student

	// التأكد أن الطالب موجود
	if err := config.DB.First(&student, id).Error; err != nil {

		// لو الطالب غير موجود
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("student not found")
		}

		return nil, err
	}

	// لو المستخدم أرسل TeacherID جديد
	if input.TeacherID != 0 {

		// نتأكد أن المدرس الجديد موجود
		var teacher models.Teacher

		if err := config.DB.First(&teacher, input.TeacherID).Error; err != nil {

			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errors.New("teacher not found")
			}

			return nil, err
		}

		// تحديث TeacherID
		student.TeacherID = input.TeacherID
	}

	// تحديث الاسم إذا تم إرساله
	if input.Name != "" {
		student.Name = input.Name
	}

	// تحديث العمر إذا تم إرساله
	if input.Age != 0 {
		student.Age = input.Age
	}

	// حفظ التعديلات
	if err := config.DB.Save(&student).Error; err != nil {
		return nil, err
	}

	// إرجاع البيانات بعد التحديث
	return &student, nil
}
