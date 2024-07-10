package service

import (
	"st/backend/db/repository"
	"st/backend/settings"
	"st/backend/utils"
)

var UserServ *UserService = &UserService{}

func InitServices() {
	UserServ.UserRepo = repository.UserRepo
	UserServ.Encrypter = &utils.Encrypter{Secret: settings.Utils.Secret}
}
