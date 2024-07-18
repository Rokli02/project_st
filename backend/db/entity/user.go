package entity

import "st/backend/db"

type User struct {
	Id        int64 `db_constraint:"PRIMARY KEY"`
	Name      *string
	Login     string `db_constraint:"UNIQUE"`
	Password  string
	DBPath    string
	CreatedAt string `db_constraint:"DEFAULT CURRENT_TIMESTAMP"`
}

var _ db.Model = (*User)(nil)

func (s *User) TableTemplate() (string, error) {
	return generateTableTemplate(*s)
}

func (*User) Migrations() []db.Migration {
	return []db.Migration{
		{
			Version:  0,
			Template: "",
		},
	}
}

func (*User) TableVersion() uint {
	return 0
}
