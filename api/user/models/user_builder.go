package models

import "strings"

type UserBuilder interface {
	WithID(id int64) UserBuilder
	WithName(name string) UserBuilder
	WithCpf(cpf string) UserBuilder
	WithEmail(email string) UserBuilder
	WithPassword(password string) UserBuilder
	Build() *User
}

type userBuilder struct {
	user *User
}

func NewBuilder() UserBuilder {
	return &userBuilder{user: &User{}}
}

func (u *userBuilder) WithID(id int64) UserBuilder {
	u.user.id = id
	return u
}

func (u *userBuilder) WithName(name string) UserBuilder {
	u.user.name = name
	return u
}

func (u *userBuilder) WithCpf(cpf string) UserBuilder {
	cpf = strings.ReplaceAll(cpf, ".", "")
	cpf = strings.ReplaceAll(cpf, "-", "")
	u.user.cpf = cpf
	return u
}

func (u *userBuilder) WithEmail(email string) UserBuilder {
	u.user.email = email
	return u
}

func (u *userBuilder) WithPassword(password string) UserBuilder {
	u.user.password = password
	return u
}

func (u *userBuilder) Build() *User {
	return u.user
}
