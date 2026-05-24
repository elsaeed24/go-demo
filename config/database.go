package config

import (
	"fmt"
	"log"
	"os"

	"go-demo/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB هو الـ instance الـ global للـ database connection
// بيتستخدم في كل الـ handlers و services
var DB *gorm.DB

// ConnectDB بتعمل connection بالـ PostgreSQL database
// بتاخد الـ credentials من الـ environment variables عشان الأمان
func ConnectDB() {
	// بنبني الـ DSN (Data Source Name) من الـ env variables
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSLMODE"),
	)

	db, err := gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{},
	)

	if err != nil {
		log.Fatal("❌ Database connection failed: ", err)
	}

	DB = db
	log.Println("✅ Database connected successfully")

	// AutoMigrate بتعمل create أو update للـ tables تلقائياً
	// لو الـ table مش موجود بيعمله، لو موجود بيعمل update بس (مش بيمسح داتا)
	DB.AutoMigrate(
		&models.Admin{},
		&models.Teacher{},
		&models.Student{},
	)

	log.Println("✅ Database migrations completed")
}
