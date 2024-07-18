package entity

import "st/backend/db"

type Metadata struct {
	Id        int64  `db_constraint:"PRIMARY KEY"`
	Key       string `db_constraint:"UNIQUE"`
	Value     *string
	Type      string `db_constraint:"DEFAULT 'app'"`
	UpdatedAt string `db_constraint:"DEFAULT CURRENT_TIMESTAMP"`
	ExpireAt  *string
}

var _ db.Model = (*Metadata)(nil)

func (s *Metadata) TableTemplate() (string, error) {
	return generateTableTemplate(*s)
}

func (*Metadata) Migrations() []db.Migration {
	return nil
}

func (*Metadata) TableVersion() uint {
	return 0
}
