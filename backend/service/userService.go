package service

import (
	"fmt"
	"st/backend/db/repository"
	"st/backend/logger"
	"st/backend/model"
	"st/backend/utils"
)

type UserService struct {
	UserRepo  *repository.UserRepository
	Encrypter *utils.Encrypter
}

var _ Service = (*UserService)(nil)

func (s *UserService) Login(user *model.LoginUser) (string, error) {
	if user == nil {
		return "", fmt.Errorf("no user is given")
	}

	// Encrypt password
	user.Password = s.Encrypter.Hash(user.Password)

	// Try to login
	userFromDB := s.UserRepo.FindOneByLoginAndPassword(user.Login, user.Password)
	// If unsuccesful return with an error
	if userFromDB == nil {
		return "", fmt.Errorf("can't log in, login name or password is incorrect")
	}

	logger.InfoF("User '%s' is logged in", userFromDB.Login)

	// Else return with the userDbPath
	return userFromDB.DBPath, nil
}

func (s *UserService) SignUp(user *model.SignUpUser) error {
	if user == nil {
		return fmt.Errorf("no user is given")
	}

	if user.Login == "" || user.Password == "" {
		return fmt.Errorf("at least one required property is missing from user")
	}

	// Check if login name is already taken
	isExist := s.UserRepo.IsExist(user.Login)
	if isExist {
		return fmt.Errorf("login name is already in use, try a different one")
	}

	// Encrypt password
	user.Password = s.Encrypter.Hash(user.Password)

	// Do something evil-ish ╰(*°▽°*)╯
	// if !isExist {
	// 	return fmt.Errorf("try again a bit later")
	// }

	result := s.UserRepo.Save(user.ToEntity())
	if !result {
		return fmt.Errorf("can sign up for some reason ㄟ( ▔, ▔ )ㄏ")
	}

	return nil
}
