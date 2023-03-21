package domain

import (
	"fmt"
	"strings"
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

func (r *repositoryMock) Create(user *models.User) error {
	if _, err := r.GetByLogin(user.GetEmail()); err == nil {
		return internal_errors.NewErrorWithCode(409, fmt.Errorf("user already exists"))
	}
	r.data[user] = true
	return nil
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

func TestLoginFail(t *testing.T) {
	data := make(map[*models.User]bool)
	repo := newMock(data)
	userDomain := New(repo)

	email := "user@example.com"
	password := "password"

	_, err := userDomain.Login(email, password)
	if err == nil {
		t.Errorf("expected error, got %v", err)
	}

	if !strings.Contains(err.Error(), "404") {
		t.Errorf("expected error 404 got %v", err)
	}
}

func TestCreateUser(t *testing.T) {
	data := make(map[*models.User]bool)
	repo := newMock(data)
	userDomain := New(repo)


	newUser := &models.UserRequest{
		Name:     "Han Solo",
		CPF:      "714.698.840-62",
		Email:    "hansolo12pc@gmail.com",
		Password: "princesslea",
	}

	user, err := userDomain.CreateUser(newUser)
	if err != nil {
		t.Errorf("user creation failed: %v", err)
	}
	if user == nil {
		t.Errorf("user creation: user is nil")
	}
}

func TestCreateUserFail(t *testing.T) {
	data := make(map[*models.User]bool)
	repo := newMock(data)
	userDomain := New(repo)

	name := "Han Solo"
	cpf := "714.698.840-62"
	email := "hansolo12pc@gmail.com"
	password := "princesslea"

	data[models.NewBuilder().WithName(name).WithCpf(cpf).WithEmail(email).WithPassword(password).Build()] = true

	newUser := &models.UserRequest{
		Name:     name,
		CPF:      cpf,
		Email:    email,
		Password: password,
	}

	_, err := userDomain.CreateUser(newUser)
	if err == nil {
		t.Errorf("expected error, got %v", err)
	}
	if !strings.Contains(err.Error(), "409") {
		t.Errorf("expected error 409 got %v", err)
	}
}

