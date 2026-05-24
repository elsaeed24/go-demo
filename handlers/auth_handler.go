package handlers

import (
	"net/http"

	"go-demo/services"

	"github.com/gin-gonic/gin"
)

// Register بتعمل admin account جديد
// POST /auth/register
// Body: { "username": "...", "password": "..." }
func Register(c *gin.Context) {
	var input services.AdminRegisterInput

	// نتحقق من الـ input (binding بيعمل validate تلقائياً)
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// نعمل الـ admin في الـ database
	admin, err := services.RegisterAdmin(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Admin created successfully",
		"admin": gin.H{
			"id":       admin.ID,
			"username": admin.Username,
		},
	})
}

// Login بتعمل authenticate للـ admin وترجعله JWT token
// POST /auth/login
// Body: { "username": "...", "password": "..." }
func Login(c *gin.Context) {
	var input services.AdminLoginInput

	// نتحقق من الـ input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// نعمل login ونجيب الـ token
	token, err := services.LoginAdmin(input)
	if err != nil {
		// 401 Unauthorized للـ credentials الغلط
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token, // الـ token ده بعته في كل request جاي
	})
}

// Me بترجع معلومات الـ admin اللي عامل login دلوقتي
// GET /auth/me  (protected - محتاج token)
func Me(c *gin.Context) {
	// نجيب البيانات اللي الـ middleware حطها في الـ context
	adminID, _ := c.Get("admin_id")
	username, _ := c.Get("admin_username")

	c.JSON(http.StatusOK, gin.H{
		"admin_id": adminID,
		"username": username,
	})
}
