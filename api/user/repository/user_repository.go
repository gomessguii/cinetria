package repository

import "github.com/gomessguii/cinetria/api/user/models"


type IUserRepository interface {
	GetAll() []*models.User
	GetByLogin(login string) (*models.User, error)
	Create(user *models.User) error
}