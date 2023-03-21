package domain

import (
	"fmt"
	"testing"

	"github.com/gomessguii/cinetria/api/user/models"
	internal_errors "github.com/gomessguii/cinetria/internal/errors"
)

type repositoryMock struct {
	data map[*models.User]bool
}

func newMock(data map[*models.User]bool) IUserRepository {
	return &repositoryMock{data}
}

func (r *repositoryMock) Create(user *models.User) (*models.User, error) {
	r.data[user] = true
	return user, nil
}

func (r *repositoryMock) GetByLogin(login string) (*models.User, error) {
	for u, _ := range r.data {
		if u.GetEmail() == login {
			return u, nil
		}
	}
	return nil, internal_errors.NewErrorWithCode(404, fmt.Errorf("user %s not found", login))
}

func TestLogin(t *testing.T) {
	data := make(map[*models.User]bool)
	repo := newMock(data)
	userDomain := New(repo)

	email := "user@example.com"
	password := "password"
	
	data[models.NewBuilder().WithEmail(email).WithPassword(password).Build()] = true

	user, err := userDomain.Login(email, password)
	if err != nil {
		t.Errorf("login failed: %v", err)
	}
	if user == nil {
		t.Errorf("login failed: user is nil")
	}
}
