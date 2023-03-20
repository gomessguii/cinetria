package domain

import "github.com/gomessguii/cinetria/api/user/models"

type IUserRepository interface {
	GetByLogin(login string) (*models.User, error)
	Create(user *models.User) (*models.User, error)
}