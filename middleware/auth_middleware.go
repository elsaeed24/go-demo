package middleware

import (
	"net/http"
	"strings"

	"go-demo/services"

	"github.com/gin-gonic/gin"
)

// AuthRequired هو الـ middleware بتاع الـ authentication
// بيتحقق إن الـ request فيه JWT token صحيح قبل ما يوصل للـ handler
// الاستخدام: router.Use(middleware.AuthRequired()) أو router.GET("/path", middleware.AuthRequired(), handler)
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// الـ token بيجي في الـ Authorization header بالشكل ده:
		// Authorization: Bearer <token>
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization header is required",
			})
			c.Abort() // وقّف الـ request هنا، متروحش للـ handler
			return
		}

		// نقسم الـ header على "Bearer " عشان ناخد الـ token بس
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization format must be: Bearer <token>",
			})
			c.Abort()
			return
		}

		//parts[0]="Bearer"
		//parts[1]="abc123456"

		tokenString := parts[1]

		// نتحقق من صحة الـ token
		claims, err := services.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid or expired token",
			})
			c.Abort()
			return
		}

		// نحفظ معلومات الـ admin في الـ context عشان الـ handlers تقدر تستخدمها
		c.Set("admin_id", claims.AdminID)
		c.Set("admin_username", claims.Username)

		// كل حاجة تمام، كمّل للـ handler
		c.Next()
	}
}
