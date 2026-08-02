package services

import (
	"errors"
	"first-app/internal/modules/user/models"
	"first-app/internal/modules/user/repositories"
	"first-app/internal/modules/user/requests/auth"
	"first-app/internal/modules/user/respones"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepository repositories.UserRepositoryInterface
}

func New() *UserService {
	return &UserService{
		userRepository: repositories.New(),
	}
}

func (userService *UserService) Creat(requset auth.RegisterRequest) (respones.User, error) {
	var response respones.User
	var user models.User

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(requset.Password), 12)
	if err != nil {
		return response, errors.New("error hash")
	}

	user.Name = requset.Name
	user.Email = requset.Email
	user.Password = string(hashedPassword)

	newUser := userService.userRepository.Create(user)
	if newUser.ID == 0 {
		return response, errors.New("error creat user")
	}
	return respones.ToUser(newUser), nil
}
