package usecase

import (
	"gotodo-app/model"
	"gotodo-app/repository"
)

type AuthorUseCase interface {
	FindAllAuthor(author string) ([]model.Author, error)
	FindAuthorByID(id string) (model.Author, error)
	FindAuthorByEmail(email string) (model.Author, error)
}

type authorUseCase struct {
	repo repository.AuthorRepository
}

func (a *authorUseCase) FindAuthorByEmail(email string) (model.Author, error) {
	return a.repo.GetByEmail(email)
}

func (a *authorUseCase) FindAllAuthor(author string) ([]model.Author, error) {
	return a.repo.List(author)
}

func (a *authorUseCase) FindAuthorByID(id string) (model.Author, error) {
	return a.repo.Get(id)
}

func NewAuthorUseCase(repo repository.AuthorRepository) AuthorUseCase {
	return &authorUseCase{repo: repo}
}
