package controllers

import (
	"net/http"
	"time"

	"example.com/v1/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Login ...
func Login(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	err := models.GetUser(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"user":    &user,
			"status":  http.StatusBadRequest,
			"message": "No User",
		})
	} else {
		sign := jwt.New(jwt.SigningMethodHS256)
		claims := sign.Claims.(jwt.MapClaims)
		claims["user"] = user.Username
		claims["exp"] = time.Now().Add(time.Minute * 15).Unix()
		token, err := sign.SignedString([]byte("secret"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Internal Serve Error",
			})
			c.Abort()
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "Login user berhasil",
			"token":   token,
		})
	}
}

// Register ...
func Register(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	err := models.RegisterUser(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	} else {
		c.JSON(http.StatusOK, user)
	}
}
