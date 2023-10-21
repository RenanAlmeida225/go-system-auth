package router

import "github.com/gin-gonic/gin"

func Initialize() {
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "hello, world",
		})
	})
	r.Run()
}
