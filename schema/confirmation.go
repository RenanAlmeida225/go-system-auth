package schema

import "time"

type Confirmation struct {
	ID        uint   `gorm:"primarykey"`
	Token     string `gorm:"unique"`
	Email     string
	CreatedAt time.Time
}
