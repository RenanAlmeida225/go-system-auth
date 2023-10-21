package schema

import "time"

type User struct {
	ID        uint `gorm:"primarykey"`
	Username  string
	Email     string `gorm:"unique"`
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
