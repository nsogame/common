package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model

	Username       string `gorm:"not null,unique"`
	UsernameCase   string `gorm:"not null,unique"`
	Email          string `gorm:"not null,unique"`
	EmailConfirmed bool
	Password       string `gorm:"not null"`
	OsuPassword    string `gorm:"not null"`

	InGameRole int

	JoinTime   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	LastActive time.Time
}
