package domain

import (
	"fmt"

	"github.com/gomessguii/cinetria/api/user/models"
	internal_errors "github.com/gomessguii/cinetria/internal/errors"
)

type IUserDomain interface {
	IsLoggedIn() bool
	Login(email, password string) (*models.User, error)
	Logoff()
	CreateUser(userRequest *models.UserRequest) (*models.User, error)
}

type userDomain struct {
	loggedUser *models.User
	repo       IUserRepository
}

func New(userRepository IUserRepository) IUserDomain {
	return &userDomain{repo: userRepository}
}


func (*userDomain) CreateUser(userRequest *models.UserRequest) (*models.User, error) {
	if !userRequest.IsFilledUp() {
		return nil, internal_errors.NewErrorWithCode(422, fmt.Errorf("CPF is invalid"))
	}

	user := userRequest.ToUser()
	
	if !user.IsCpfValid() {
		return nil, internal_errors.NewErrorWithCode(422, fmt.Errorf("CPF is invalid"))
	}

	return user, nil
}

func (u *userDomain) IsLoggedIn() bool {
	return !u.loggedUser.IsNil()
}

func (d *userDomain) Login(email, password string) (*models.User, error) {
	user, err := d.repo.GetByLogin(email)

	if err != nil {
		return nil, err
	}

	if !user.IsPasswordCorrect(password) {
		return nil, fmt.Errorf("password is incorrect")
	}

	d.loggedUser = user
	return user, nil
}

func (d *userDomain) Logoff() {
	d.loggedUser = nil
}
