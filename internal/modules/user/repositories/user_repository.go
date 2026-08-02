package repositories

import (
	"first-app/internal/modules/user/models"
	"first-app/pkg/database"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func New() *UserRepository {
	return &UserRepository{
		DB: database.Connections(),
	}
}

func (userrepository *UserRepository) Create(user models.User) models.User {
	var newUser models.User
	userrepository.DB.Create(&user).Scan(&newUser)
	return newUser
}
