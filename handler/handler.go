package handler

import (
	"github.com/RenanAlmeida225/go-system-auth/config"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func InitializeHandler() {
	db = config.GetPostgres()
}
