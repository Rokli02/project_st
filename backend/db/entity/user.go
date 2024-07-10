package entity

import "st/backend/db"

type User struct {
	Id        int64
	Name      string
	Login     string
	Password  string
	DBPath    string
	CreatedAt string
}

var _ db.Model = (*User)(nil)

var UserTableTemplate string = `CREATE TABLE %s (
	id INTEGER PRIMARY KEY,
	name TEXT,
	login TEXT NOT NULL UNIQUE,
	password TEXT NOT NULL,
	dbPath TEXT
);`

var UserTableVersion uint = 0

var UserTableMigrations = []db.Migration{
	{
		Version:  1,
		Template: "",
	},
}
