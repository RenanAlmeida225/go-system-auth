package handler

import (
	"fmt"
	"net/http"

	"github.com/RenanAlmeida225/go-system-auth/helper"
	"github.com/RenanAlmeida225/go-system-auth/schema"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(ctx *gin.Context) {
	request := SignUpRequest{}
	ctx.BindJSON(&request)
	if err := request.Validate(); err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user := schema.User{}
	if err := db.First(&user, "email=?", request.Email).Error; err == nil {
		sendError(ctx, http.StatusBadRequest, "email invalid")
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.MinCost)
	if err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	token := uuid.New().String()

	user.Username = request.Username
	user.Email = request.Email
	user.Password = string(passwordHash)
	user.Confirmation = schema.Confirmation{
		Token: token,
		Email: user.Email,
	}

	if err = db.Create(&user).Error; err != nil {
		fmt.Println(err)
		sendError(ctx, http.StatusBadRequest, "fail on register user")
		return
	}

	if err = helper.SendEmail(request.Email, token); err != nil {
		sendError(ctx, http.StatusBadRequest, "fail on send email")
		return
	}

	send(ctx, http.StatusCreated, "sign up", "register user successfully")

}
