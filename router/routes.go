package router

import (
	"github.com/RenanAlmeida225/go-system-auth/handler"
	"github.com/RenanAlmeida225/go-system-auth/middleware"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(r *gin.Engine) {
	handler.InitializeHandler()
	v1 := r.Group("/api/v1")
	auth := v1.Group("auth")
	{
		auth.POST("/signUp", handler.SignUp)
		auth.GET("/enable/:token", handler.EnableUser)
		auth.POST("/signIn", handler.SignIn)
	}
	products := v1.Group("products", middleware.VerifyAuth)
	{
		products.GET("/save")
	}
}
