package routes

import (
	"net/http"

	"go-demo/handlers"
	"go-demo/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRoutes بتعمل تنظيم كل الـ routes في البروجكت
// بتقسمهم لـ public routes (مش محتاجة token) وprotected routes (محتاجة token)
func SetupRoutes(router *gin.Engine) {

	// ===== Health Check =====
	// مش محتاج auth عشان تقدر تتأكد إن الـ server شغال
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Server is running ✅",
		})
	})

	// ===== Auth Routes (Public - مش محتاجة token) =====
	auth := router.Group("/auth")
	{
		auth.POST("/register", handlers.Register) // إنشاء admin جديد
		auth.POST("/login", handlers.Login)       // الـ login والحصول على token
	}

	// ===== Protected Routes (محتاجة JWT token) =====
	// كل الـ routes الجاية دي محتاجة Authorization: Bearer <token>
	protected := router.Group("/")
	protected.Use(middleware.AuthRequired()) // بيطبق الـ middleware على كل الـ routes الجوه
	{
		// Auth
		protected.GET("/auth/me", handlers.Me) // معلومات الـ admin اللي عامل login

		// Students
		protected.POST("/students", handlers.CreateStudent) // إضافة student
		protected.GET("/students", handlers.GetStudents)    // جلب كل الـ students

		// Teachers
		protected.POST("/teachers", handlers.CreateTeacher) // إضافة teacher
		protected.GET("/teachers", handlers.GetTeachers)    // جلب كل الـ teachers
	}
}
