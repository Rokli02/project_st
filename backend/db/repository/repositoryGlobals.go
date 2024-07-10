package repository

var UserRepo *UserRepository = &UserRepository{}
var MetadataRepo *MetadataRepository = &MetadataRepository{}

func InitRepositories() {
	UserRepo.modelName = "user"
	MetadataRepo.modelName = "metadata"
}
