package services

import (
	"first-app/internal/modules/user/requests/auth"
	"first-app/internal/modules/user/respones"
)

type UserServiceInterface interface {
	Creat(requset auth.RegisterRequest) (respones.User,error)

}
