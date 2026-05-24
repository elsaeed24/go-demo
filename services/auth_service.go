package services

import (
	"errors"
	"os"
	"strconv"
	"time"

	"go-demo/config"
	"go-demo/models"

	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt/v5"
)

// AdminLoginInput هو الـ struct اللي بياخد الـ data من الـ request
type AdminLoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// AdminRegisterInput هو الـ struct اللي بيتستخدم في إنشاء admin جديد
type AdminRegisterInput struct {
	Username string `json:"username" binding:"required,min=3"`
	Password string `json:"password" binding:"required,min=6"`
}

// Claims هو الـ struct اللي بيتخزن في الـ JWT token
// بيحتوي على معلومات الـ admin زي الـ ID و Username
type Claims struct {
	AdminID  uint   `json:"admin_id"`
	Username string `json:"username"`

	// بيورث الـ standard JWT fields زي ExpiresAt و IssuedAt
	jwt.RegisteredClaims
}

// HashPassword بتحول الـ plain text password لـ bcrypt hash آمن
// cost = 12 يعني صعب على الـ brute force attacks
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(bytes), err
}

// CheckPassword بتقارن الـ plain text password بالـ hashed password
// بترجع true لو الـ password صح
func CheckPassword(hashedPassword, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}

// GenerateToken بتعمل JWT token جديد للـ admin بعد الـ login
// الـ token بيتشفر بالـ secret key الموجود في الـ .env
func GenerateToken(admin models.Admin) (string, error) {
	// نجيب مدة صلاحية الـ token من الـ .env
	hoursStr := os.Getenv("JWT_EXPIRES_HOURS")
	hours, err := strconv.Atoi(hoursStr)
	if err != nil {
		hours = 24 // default 24 ساعة لو مفيش قيمة في الـ .env
	}

	// نعمل الـ claims (البيانات اللي جوه الـ token)
	claims := &Claims{
		AdminID:  admin.ID,
		Username: admin.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			// الـ token بيبطل بعد المدة المحددة
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(hours) * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// بنعمل token بـ HS256 algorithm (HMAC SHA-256)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// بنوقّع الـ token بالـ secret key
	secret := os.Getenv("JWT_SECRET")
	return token.SignedString([]byte(secret))
}

// ValidateToken بتتحقق إن الـ JWT token صحيح وصالح
// بترجع الـ claims لو الـ token تمام
func ValidateToken(tokenString string) (*Claims, error) {
	secret := os.Getenv("JWT_SECRET")

	claims := &Claims{}

	// بنعمل parse للـ token ونتحقق من الـ signature
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// نتأكد إن الـ algorithm هو HS256 بالظبط
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid or expired token")
	}

	return claims, nil
}

// RegisterAdmin بتعمل admin جديد في الـ database
// بتـ hash الـ password قبل ما يتخزن
func RegisterAdmin(input AdminRegisterInput) (*models.Admin, error) {
	// نتأكد إن الـ username مش موجود قبل كده
	var existing models.Admin
	if err := config.DB.Where("username = ?", input.Username).First(&existing).Error; err == nil {
		return nil, errors.New("username already exists")
	}

	// نعمل hash للـ password
	hashedPassword, err := HashPassword(input.Password)
	if err != nil {
		return nil, errors.New("failed to process password")
	}

	// نعمل الـ admin في الـ database
	admin := models.Admin{
		Username: input.Username,
		Password: hashedPassword,
	}

	if err := config.DB.Create(&admin).Error; err != nil {
		return nil, errors.New("failed to create admin")
	}

	return &admin, nil
}

// LoginAdmin بتتحقق من الـ credentials وترجع JWT token
func LoginAdmin(input AdminLoginInput) (string, error) {
	var admin models.Admin

	// نبحث عن الـ admin بالـ username
	if err := config.DB.Where("username = ?", input.Username).First(&admin).Error; err != nil {
		// مهم: نرجع نفس الـ error للـ username والـ password عشان الأمان
		return "", errors.New("invalid username or password")
	}

	// نتحقق من الـ password
	if !CheckPassword(admin.Password, input.Password) {
		return "", errors.New("invalid username or password")
	}

	// نعمل JWT token
	token, err := GenerateToken(admin)
	if err != nil {
		return "", errors.New("failed to generate token")
	}

	return token, nil
}
