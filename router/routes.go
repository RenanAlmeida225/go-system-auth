package router

import (
	"github.com/RenanAlmeida225/go-system-auth/handler"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(r *gin.Engine) {
	handler.InitializeHandler()
	v1 := r.Group("/api/v1")
	auth := v1.Group("auth")
	{
		auth.POST("/signUp", handler.SignUp)
		auth.POST("/signIn")
	}
}
