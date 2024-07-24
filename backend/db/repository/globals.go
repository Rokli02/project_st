package repository

import "st/backend/db/entity"

var User *UserRepository = &UserRepository{}
var Metadata *MetadataRepository = &MetadataRepository{}

func InitRepositories() {
	User.modelName = entity.NameOfModel(entity.User{})
	Metadata.modelName = entity.NameOfModel(entity.Metadata{})
}
