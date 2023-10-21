package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Init() error {
	var err error
	if err = InitializeGodotenv(); err != nil {
		return fmt.Errorf("error on load environment variables")
	}

	db, err = InitializePostgres()
	if err != nil {
		return fmt.Errorf("error on init postgres")
	}

	return nil
}

func GetPostgres() *gorm.DB {
	return db
}
