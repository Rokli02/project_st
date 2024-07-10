package model

import "st/backend/db/entity"

type User struct {
	Id    int64
	Name  string
	Login string
}

type LoginUser struct {
	Login    string
	Password string
}

type SignUpUser struct {
	Login    string
	Name     string
	Password string
}

func (u *User) FromEntity(user *entity.User) {
	if u == nil {
		u = &User{}
	}

	u.Id = user.Id
	u.Login = user.Login
	u.Name = user.Name
}

func (u *LoginUser) ToEntity() *entity.User {
	return &entity.User{
		Login:    u.Login,
		Password: u.Password,
	}
}

func (u *SignUpUser) ToEntity() *entity.User {
	return &entity.User{
		Login:    u.Login,
		Name:     u.Name,
		Password: u.Password,
	}
}
