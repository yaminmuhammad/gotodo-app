package repo_mock

import (
	"gotodo-app/model"
	"gotodo-app/shared/shared_model"

	"github.com/stretchr/testify/mock"
)

type TaskRepoMock struct {
	mock.Mock
}

func (t *TaskRepoMock) List(page, size int) ([]model.Task, shared_model.Paging, error) {
	args := t.Called(page, size)
	return args.Get(0).([]model.Task), args.Get(1).(shared_model.Paging), args.Error(2)
}
func (t *TaskRepoMock) Create(payload model.Task) (model.Task, error) {
	args := t.Called(payload)
	return args.Get(0).(model.Task), args.Error(1)
}
func (t *TaskRepoMock) GetByAuthor(authorId string) ([]model.Task, error) {
	args := t.Called(authorId)
	return args.Get(0).([]model.Task), args.Error(2)
}
