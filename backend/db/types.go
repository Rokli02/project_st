package db

import "database/sql"

type ConnectTypes uint8

const (
	CONNECT_IF_EXISTS        ConnectTypes = iota // Connects to database only if it was created previously (already exists)
	CREATE_IF_NEEDED                             // Connects to database. If it wasn't existing yet, creates it.
	CREATE_NEW_IF_NOT_EXISTS                     // Creates a new database, only if it wasn't existing previously.
	CREATE_ALWAYS                                // Recreates the Database everytime with its fresh table schemas
)

// Can be used to migrate up from a previous table version
type Migration struct {
	Version  uint   // Version number starting from '0' and incrementing by '1' everytime
	Template string // SQL Command that increments the previous Table version to the next one
}

type Model interface {
}

type Repository interface {
	SetDB(*sql.DB)
	ModelName() string
	CreateTable() bool
	IsTableExist() bool
	Migrate() uint
	DropTableTemplate() bool
}
