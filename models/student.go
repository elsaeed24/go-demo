package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model

	Name string
	Age  int

	TeacherID uint
	Teacher   Teacher
}
