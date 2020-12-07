package controllers

import (
	"net/http"
	"time"

	"example.com/v1/models"
	"example.com/v1/validation"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

// Login ...
func Login(c *gin.Context) {
	var user models.User
	var url string = "localhost:8080/api/v1/todo"
	c.BindJSON(&user)
	err := models.GetUser(&user)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"user":    &user,
			"status":  http.StatusBadRequest,
			"message": "No User",
		})
	} else {
		sign := jwt.New(jwt.SigningMethodHS256)
		claims := sign.Claims.(jwt.MapClaims)
		claims["user"] = user.Username
		claims["exp"] = time.Now().Add(time.Minute * 1).Unix()
		token, err := sign.SignedString([]byte("secret"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Internal Serve Error",
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "Login user berhasil",
			"result":  token,
		})
		c.Redirect(http.StatusMovedPermanently, url)
	}
}

// Register ...
func Register(c *gin.Context) {
	var user models.User
	c.BindJSON(&user)
	validate := validation.RegisterUser(&user)
	if validate != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": validate,
		})
	} else {
		err := models.RegisterUser(&user)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Internal Serve Error",
			})
		} else {
			result := make(map[string]interface{})
			result["test"] = "percobaan"
			mapstructure.Decode(&user, &result)
			c.JSON(http.StatusOK, gin.H{
				"status":  http.StatusOK,
				"message": "Register user berhasil",
				"result":  result,
			})
		}
	}

}

// LoginForm ...
func LoginForm(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Silahkan Login",
	})
}

// RegisterForm ...
func RegisterForm(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Register Form",
	})
}
