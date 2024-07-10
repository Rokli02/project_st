package entity

import "st/backend/db"

type Metadata struct {
	Id        int64
	Key       string
	Value     string
	Type      string
	UpdatedAt string
	ExpireAt  string
}

var _ db.Model = (*Metadata)(nil)

var MetadataTableTemplate string = `CREATE TABLE %s (
	id INTEGER PRIMARY KEY,
	Key Text NOT NULL,
	Value Text NOT NULL,
	Type Text DEFAULT 'app',
	UpdatedAt Text NOT NULL DEFAULT CURRENT_TIMESTAMP,
	ExpireAt Text DEFAULT NULL
);`

var MetadataTableVersion uint = 0

var MetadataTableMigrations = []db.Migration{
	{
		Version:  1,
		Template: "",
	},
}
