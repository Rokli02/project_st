package service

import (
	"st/backend/db/repository"
	"st/backend/utils"
	"st/backend/utils/settings"
)

var User *UserService = &UserService{}
var Metadata *MetadataService = &MetadataService{}

func InitServices() {
	User.UserRepo = repository.User
	User.Encrypter = utils.NewEncrypter(settings.Utils.Secret)

	Metadata.MetaRepo = repository.Metadata
}
