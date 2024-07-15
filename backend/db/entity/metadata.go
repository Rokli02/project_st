package entity

import "st/backend/db"

type Metadata struct {
	Id        int64  `db_constraint:"PRIMARY KEY"`
	Key       string `db_constraint:"NOT NULL"`
	Value     string `db_constraint:"NOT NULL"`
	Type      string `db_constraint:"DEFAULT 'app'"`
	UpdatedAt string `db_constraint:"DEFAULT CURRENT_TIMESTAMP"`
	ExpireAt  string `db_constraint:"DEFAULT NULL"`
}

var _ db.Model = (*Metadata)(nil)

func (s *Metadata) TableTemplate() (string, error) {
	return generateTableTemplate(*s)
}

func (*Metadata) Migrations() []db.Migration {
	return metadataTableMigrations
}

var MetadataTableVersion uint = 0

var metadataTableMigrations = []db.Migration{
	{
		Version:  1,
		Template: "",
	},
}
