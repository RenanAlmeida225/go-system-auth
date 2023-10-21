package middleware

import (
	"net/http"
	"strings"

	"github.com/RenanAlmeida225/go-system-auth/helper"
	"github.com/gin-gonic/gin"
)

func VerifyAuth(ctx *gin.Context) {
	header := ctx.GetHeader("Authorization")
	if header == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "bad header value given",
		})
		return
	}

	jwtToken := strings.Split(header, " ")
	if len(jwtToken) != 2 {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "incorrectly formatted authorization header",
		})
		return
	}
	if err := helper.VerifyJWT(jwtToken[1]); err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Token invalid",
		})
		return
	}
	ctx.Next()
}
