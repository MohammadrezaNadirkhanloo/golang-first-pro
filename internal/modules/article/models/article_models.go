package models

import (
	"first-app/internal/modules/user/models"

	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title string
	Content string
	UserID uint
	User models.User
}
