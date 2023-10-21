package router

import "github.com/gin-gonic/gin"

func InitializeRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	auth := v1.Group("auth")
	{
		auth.POST("/signUp")
		auth.POST("/signIn")
	}
}
