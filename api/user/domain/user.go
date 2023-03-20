package domain

import "github.com/gomessguii/cinetria/api/user/models"

type IUserDomain interface {
	CreateUser(userRequest *models.UserRequest) (*models.User, error)
}
