package settings

type DatabaseSettings struct {
	BaseDatabaseName string
}

type RepositorySettings struct{}

type ServiceSettings struct{}

type UtilsSettings struct {
	Secret string
}

type MetadataKeysConfig struct {
	CurrentUserId string
}
