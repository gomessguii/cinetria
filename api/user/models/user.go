package models

type User struct {
	id       int64
	name     string
	cpf      string
	email    string
	password string
}

func (u *User) GetEmail() string {
	return u.email
}

func (u *User) IsPasswordCorrect(password string) bool {
	return u.password == password
}

func (u *User) IsCpfValid() bool {
	if len(u.cpf) != 11 {
		return false
	}

	digits := make([]int, 11)

	for i := 0; i < 9; i++ {
		digits[i] = int(u.cpf[i] - '0')
	}

	digits[9], digits[10] = 0, 0

	for i := 0; i < 9; i++ {
		digits[9] += digits[i] * (10 - i)
		digits[10] += digits[i] * (11 - i)
	}

	digits[9] %= 11

	if digits[9] < 2 {
		digits[9] = 0
	} else {
		digits[9] = 11 - digits[9]
	}

	digits[10] += digits[9] * 2

	digits[10] %= 11

	if digits[10] < 2 {
		digits[10] = 0
	} else {
		digits[10] = 11 - digits[10]
	}

	return digits[9] == int(u.cpf[9]-'0') && digits[10] == int(u.cpf[10]-'0')
}
