package service

import (
	"fmt"
	"st/backend/db/repository"
	"st/backend/model"
	"st/backend/utils"
	"st/backend/utils/lang"
	"st/backend/utils/logger"
)

type UserService struct {
	UserRepo  *repository.UserRepository
	Encrypter *utils.Encrypter
}

var _ Service = (*UserService)(nil)

func (s *UserService) Login(user *model.LoginUser) (string, error) {
	if user == nil {
		return "", fmt.Errorf(lang.Text.User.Get("NO_USER_GIVEN"))
	}

	// Encrypt password
	user.Password = s.Encrypter.Hash(user.Password)

	// Try to login
	userFromDB := s.UserRepo.FindOneByLoginAndPassword(user.Login, user.Password)
	// If unsuccesful return with an error
	if userFromDB == nil {
		return "", fmt.Errorf(lang.Text.User.Get("INVALID_LOGIN_OR_PASSWORD"))
	}

	logger.InfoF("User '%s' is logged in", userFromDB.Login)

	// Else return with the userDbPath
	return userFromDB.DBPath, nil
}

func (s *UserService) SignUp(user *model.SignUpUser) error {
	if user == nil {
		return fmt.Errorf(lang.Text.User.Get("NO_USER_GIVEN"))
	}

	if user.Login == "" || user.Password == "" {
		return fmt.Errorf(lang.Text.Common.Get("REQUIRED_PROP_MISSING"))
	}

	// Check if login name is already taken
	isExist := s.UserRepo.IsExist(user.Login)
	if isExist {
		return fmt.Errorf(lang.Text.User.Get("LOGIN_IS_ALREADY_IN_USE"))
	}

	// Encrypt password
	user.Password = s.Encrypter.Hash(user.Password)

	// Do something evil-ish ╰(*°▽°*)╯
	// if !isExist {
	// 	return fmt.Errorf("try again a bit later")
	// }

	result := s.UserRepo.Save(user.ToEntity())
	if !result {
		return fmt.Errorf(lang.Text.User.Get("UNKNOWN_SIGN_UP_ERROR")) // USER[UNKNOWN_SIGN_UP_ERROR]
	}

	return nil
}
