package models

import "gorm.io/gorm"

// Admin هو الموديل الخاص بالمسؤولين عن النظام
// بس الـ Admin هو اللي يقدر يعمل login ويتحكم في الـ system
type Admin struct {
	gorm.Model

	// Username لازم يكون unique لكل admin
	Username string `gorm:"uniqueIndex;not null" json:"username"`

	// Password بيتخزن كـ hashed (مش plain text أبداً)
	Password string `gorm:"not null" json:"-"` // json:"-" يمنع إن الـ password يتبعت في الـ response
}
