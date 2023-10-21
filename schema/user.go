package schema

import "time"

type User struct {
	ID           uint `gorm:"primarykey"`
	Username     string
	Email        string `gorm:"unique"`
	Password     string
	IsEnabled    bool         `gorm:"default:false"`
	Confirmation Confirmation `gorm:"foreignKey:Email;references:email"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
