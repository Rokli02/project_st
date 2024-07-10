package settings

var Utils UtilsSettings = UtilsSettings{}
var MetadataKeys MetadataKeysConfig = MetadataKeysConfig{}

func InitSettings() {
	Utils.Secret = "ooo"
	MetadataKeys.CurrentUserId = "currentUserId"
}
