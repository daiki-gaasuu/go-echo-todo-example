package domain

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string
	Name     string
	Password string
	UserId   string
}
