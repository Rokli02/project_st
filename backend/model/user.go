package model

import "st/backend/db/entity"

type User struct {
	Id     int64   `json:"id"`
	Name   *string `json:"name"`
	Login  string  `json:"login"`
	DBPath *string `json:"-"`
}

type LoginUser struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type SignUpUser struct {
	Login    string  `json:"login"`
	Name     *string `json:"name"`
	Password string  `json:"password"`
}

func UserEntityToUser(entity *entity.User) *User {
	u := &User{}

	u.Id = entity.Id
	u.Login = entity.Login
	u.Name = entity.Name
	u.DBPath = &entity.DBPath

	return u
}

func (user *SignUpUser) ToEntity() *entity.User {
	return &entity.User{
		Login:    user.Login,
		Name:     user.Name,
		Password: user.Password,
	}
}
