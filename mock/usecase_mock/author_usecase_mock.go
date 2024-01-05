package usecase_mock

import (
	"gotodo-app/model"

	"github.com/stretchr/testify/mock"
)

type AuthorUseCaseMock struct {
	mock.Mock
}

func (a *AuthorUseCaseMock) FindAllAuthor(author string) ([]model.Author, error) {
	args := a.Called(author)
	return args.Get(0).([]model.Author), args.Error(1)
}

func (a *AuthorUseCaseMock) FindAuthorByID(id string) (model.Author, error) {
	args := a.Called(id)
	return args.Get(0).(model.Author), args.Error(1)
}

func (a *AuthorUseCaseMock) FindAuthorByEmail(email string) (model.Author, error) {
	args := a.Called(email)
	return args.Get(0).(model.Author), args.Error(1)
}
