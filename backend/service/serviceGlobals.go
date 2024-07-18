package service

import (
	"st/backend/db/repository"
	"st/backend/settings"
	"st/backend/utils"
)

var User *UserService = &UserService{}
var Metadata *MetadataService = &MetadataService{}

func InitServices() {
	User.UserRepo = repository.User
	User.Encrypter = &utils.Encrypter{Secret: settings.Utils.Secret}

	Metadata.MetaRepo = repository.Metadata
}
