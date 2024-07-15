package entity

import "st/backend/db"

type User struct {
	Id        int64 `db_constraint:"PRIMARY KEY"`
	Name      string
	Login     string `db_constraint:"UNIQUE NOT NULL"`
	Password  string `db_constraint:"NOT NULL"`
	DBPath    string
	CreatedAt string `db_constraint:"DEFAULT CURRENT_TIMESTAMP"`
}

var _ db.Model = (*User)(nil)

func (s *User) TableTemplate() (string, error) {
	return generateTableTemplate(*s)
}

func (*User) Migrations() []db.Migration {
	return userTableMigrations
}

var UserTableVersion uint = 0

var userTableMigrations = []db.Migration{
	{
		Version:  1,
		Template: "",
	},
}
