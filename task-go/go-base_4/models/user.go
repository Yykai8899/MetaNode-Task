package models

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"task-go/task-go/go-base_4/dao"
	"time"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Email    string
}

var jwtKey = []byte("secret_key")

func UserInitDB() *gorm.DB {
	db := dao.ConnectDB()
	err := db.AutoMigrate(&User{})
	if err != nil {
		panic(err)
	}
	return db
}

func Register(user User) (err error) {
	db := UserInitDB()
	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	if err := db.Create(&user).Error; err != nil {
		return err
	}
	return
}

func Login(c *gin.Context, user User) (string, error) {
	db := UserInitDB()
	var storedUser User
	if err := db.Where("username = ?", user.Username).First(&storedUser).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return "", err
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return "", err
	}

	// 生成 JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       storedUser.ID,
		"username": storedUser.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return "", err
	}
	fmt.Println(tokenString)
	return tokenString, nil
}

func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		var token = c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "必须上传JWT签名信息"})
		}

		// 鉴权操作
		var loginParam = jwt.MapClaims{}
		jwtToken, err := jwt.ParseWithClaims(token, &loginParam, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return jwtKey, nil // 使用相同的密钥
		})

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "无效的 token"})
			return
		}
		if loginParam, ok := jwtToken.Claims.(*jwt.MapClaims); ok && jwtToken.Valid {
			fmt.Printf("%v", loginParam)
			c.Next()
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "无效的 token"})
		}
	}
}
