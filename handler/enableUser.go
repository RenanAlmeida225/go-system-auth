package handler

import (
	"net/http"

	"github.com/RenanAlmeida225/go-system-auth/schema"
	"github.com/gin-gonic/gin"
)

func EnableUser(ctx *gin.Context) {
	token := ctx.Param("token")

	confirmation := schema.Confirmation{}
	user := &confirmation.User
	if err := db.Where("token=?", token).Preload("User").First(&confirmation).Error; err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user.IsEnabled = true
	if err := db.Save(&user).Error; err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}
	if err := db.Delete(&confirmation).Error; err != nil {
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	send(ctx, http.StatusOK, "enable user", "confirmed email")
}
