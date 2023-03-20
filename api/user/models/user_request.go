package models

type UserRequest struct {
	Name string `json:"name"`
	CPF string `json:"cpf"`
	Email string `json:"email"`
	Password string `json:"password"`
}

func (u *UserRequest)IsFilledUp() bool {
	if u == nil {
		return false
	}
	if len(u.Name) == 0 {
		return false
	}
	if len(u.CPF) == 0 {
		return false
	}
	if len(u.Email) == 0 {
		return false
	}
	if len(u.Password) == 0 {
		return false
	}
	return true
}

func (u *UserRequest) ToUser() *User {
	return NewBuilder().WithName(u.Name).WithCpf(u.CPF).WithEmail(u.Email).WithPassword(u.Password).Build()
}