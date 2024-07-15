package settings

var Database DatabaseSettings = DatabaseSettings{}
var Utils UtilsSettings = UtilsSettings{}
var MetadataKeys MetadataKeysConfig = MetadataKeysConfig{}

func InitSettings() {
	Database.BaseDatabaseName = "testR_hablagy"
	Utils.Secret = "ooo"
	MetadataKeys.CurrentUserId = "currentUserId"
}
