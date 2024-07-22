package settings

type ConnectTypes uint8

const (
	CONNECT_IF_EXISTS ConnectTypes = iota // Connects to database only if it was created previously (already exists)
	CREATE_IF_NEEDED                      // Connects to database. If it wasn't existing yet, creates it.
	CREATE_ALWAYS                         // Recreates the Database everytime with its fresh table schemas
)

type AppSettings struct {
	BaseDatabaseConnectType ConnectTypes
	Version                 string
}

type DatabaseSettings struct {
	BaseDatabaseName string
	DateFormat       string
}

type RepositorySettings struct{}

type ServiceSettings struct{}

type UtilsSettings struct {
	Secret string
}

type MetadataKeysConfig struct {
	CurrentUserId    string
	LanguageId       string
	UserTableVersion string
}
