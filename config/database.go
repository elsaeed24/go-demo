package config

import (
	"log"

	"go-demo/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {

	dsn := "host=localhost user=admin password=123456 dbname=school port=5433 sslmode=disable"

	db, err := gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{},
	)

	if err != nil {
		log.Fatal("Database Error")
	}

	DB = db

	DB.AutoMigrate(
		&models.Teacher{},
		&models.Student{},
	)
}
