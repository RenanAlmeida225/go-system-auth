package handler

import (
	"net/http"

	"github.com/RenanAlmeida225/go-system-auth/helper"
	"github.com/RenanAlmeida225/go-system-auth/schema"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SignIn(ctx *gin.Context) {
	request := SignInRequest{}
	ctx.BindJSON(&request)
	if err := request.Validate(); err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	user := schema.User{}
	if err := db.First(&user, "email=?", request.Email).Error; err != nil {
		sendError(ctx, http.StatusBadRequest, "user not found")
		return
	}

	if !user.IsEnabled {
		sendError(ctx, http.StatusUnauthorized, "confirm your email")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		sendError(ctx, http.StatusBadRequest, "password incorrect")
		return
	}

	jwt, err := helper.CreateJwt(user.Email, user.IsEnabled)
	if err != nil {
		sendError(ctx, http.StatusBadRequest, "error on created jwt")
		return
	}

	send(ctx, http.StatusOK, "sign in", jwt)
}
