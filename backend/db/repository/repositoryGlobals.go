package repository

import "st/backend/db/entity"

var UserRepo *UserRepository = &UserRepository{}
var MetadataRepo *MetadataRepository = &MetadataRepository{}

func InitRepositories() {
	UserRepo.modelName = entity.NameOfModel(entity.User{})
	MetadataRepo.modelName = entity.NameOfModel(entity.Metadata{})
}
