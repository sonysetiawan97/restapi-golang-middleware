package route

import (
	"example.com/v1/controllers"
	"example.com/v1/middleware"
	"github.com/gin-gonic/gin"
)

// Routes ...
func Routes() *gin.Engine {
	r := gin.Default()
	login := r.Group("/")
	{
		login.GET("login", controllers.LoginForm)
		login.POST("login", controllers.Login)
		login.GET("register", controllers.RegisterForm)
		login.POST("register", controllers.Register)
	}
	v1 := r.Group("/api/v1", middleware.AuthJWT)
	{
		v1.GET("todo", controllers.GetTodos)
		v1.POST("todo", controllers.CreateATodo)
		v1.GET("todo/:id", controllers.GetATodo)
		v1.PUT("todo/:id", controllers.UpdateATodo)
		v1.DELETE("todo/:id", controllers.DeleteATodo)
	}
	return r
}
