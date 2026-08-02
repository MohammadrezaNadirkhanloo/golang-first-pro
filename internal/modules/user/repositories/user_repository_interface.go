package repositories

import "first-app/internal/modules/user/models"

type UserRepositoryInterface interface {
	Create(user models.User) models.User
}
