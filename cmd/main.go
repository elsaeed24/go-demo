package main

import (
	"log"
	"os"

	"go-demo/config"
	"go-demo/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// نحمّل الـ .env file في أول حاجة
	// godotenv بيحمّل المتغيرات اللي في .env ويحطها في os.Getenv
	if err := godotenv.Load(); err != nil {
		log.Fatal("❌ Error loading .env file: ", err)
	}

	log.Println("✅ Environment variables loaded")

	// نعمل connection بالـ database
	config.ConnectDB()

	// نعمل الـ Gin router
	router := gin.Default()

	// نضيف كل الـ routes
	routes.SetupRoutes(router)

	// نجيب الـ port من الـ .env، default 8080
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 Server starting on port %s", port)

	// نشغّل الـ server
	if err := router.Run(":" + port); err != nil {
		log.Fatal("❌ Server failed to start: ", err)
	}
}
