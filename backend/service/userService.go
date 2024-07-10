package service

import (
	"fmt"
	"st/backend/db/repository"
	"st/backend/model"
	"st/backend/utils"
)

type UserService struct {
	UserRepo  *repository.UserRepository
	Encrypter *utils.Encrypter
}

var _ Service = (*UserService)(nil)

func (s *UserService) Login(user *model.LoginUser) (string, error) {
	// Encrypt password

	// Try to login

	// If succesful, return with the userDbPath

	// Else return with an error

	return "", fmt.Errorf("unable to login")
}

func (s *UserService) SignUp(user *model.SignUpUser) {
	// Check if login name is already taken

	// Encrypt password
}
