package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model // فيه ID + timestamps

	Name string
	Age  int

	TeacherID uint
	Teacher   Teacher
}
