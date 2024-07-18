package settings

var App AppSettings = AppSettings{}
var Database DatabaseSettings = DatabaseSettings{}
var Utils UtilsSettings = UtilsSettings{}
var MetadataKeys MetadataKeysConfig = MetadataKeysConfig{}

func InitSettings() {
	App.Version = "0.0.1"
	App.BaseDatabaseConnectType = CONNECT_IF_EXISTS

	Database.BaseDatabaseName = "testR_hablagy"
	Database.DateFormat = "2006-01-02T15:04:05"

	Utils.Secret = "ooo"

	MetadataKeys.CurrentUserId = "currentUserId"
	MetadataKeys.UserTableVersion = "userTableVersion"
}
