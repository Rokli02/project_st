package db

import "database/sql"

// Can be used to migrate up from a previous table version
type Migration struct {
	Version  uint   // Version number starting from '0' and incrementing by '1' everytime
	Template string // SQL Command that increments the previous Table version to the next one
}

type Model interface {
	TableTemplate() (string, error)
	// The migration list must increase strictly by version number.
	Migrations() []Migration
	TableVersion() uint
}

type ModelField struct {
	Name        string
	Type        string
	Constraints string
}

type Repository interface {
	SetDB(*sql.DB)
	ModelName() string
	CreateTable() bool
	IsTableExist() bool
	InitTable()
	Migrate() uint
	DropTable() bool
}
