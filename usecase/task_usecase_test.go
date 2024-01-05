package usecase

import (
	"fmt"
	"gotodo-app/mock/repo_mock"
	"gotodo-app/mock/usecase_mock"
	"gotodo-app/model"
	"gotodo-app/shared/shared_model"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var expectedTask = model.Task{
	ID:        "1",
	Title:     "This is title",
	Content:   "This is content",
	AuthorId:  expectedAuthor.ID,
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
}

var expectedAuthor = model.Author{
	ID:        "1",
	Name:      "This is name",
	Email:     "mail@mail.com",
	Password:  "password",
	Role:      "user",
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
	Tasks: []model.Task{
		{
			ID:        "1",
			Title:     "This is title",
			Content:   "This is content",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	},
}

type TaskUseCaseTestSuite struct {
	suite.Suite
	trm *repo_mock.TaskRepoMock
	aum *usecase_mock.AuthorUseCaseMock
	tuc TaskUseCase
}

func (suite *TaskUseCaseTestSuite) SetupTest() {
	suite.trm = new(repo_mock.TaskRepoMock)
	suite.aum = new(usecase_mock.AuthorUseCaseMock)
	suite.tuc = NewTaskUseCase(suite.trm, suite.aum)
}

func (suite *TaskUseCaseTestSuite) TestRegisterNewTask_Success() {
	suite.aum.On("FindAuthorByID", expectedTask.AuthorId).Return(expectedAuthor, nil)
	suite.trm.On("Create", expectedTask).Return(expectedTask, nil)
	actual, err := suite.tuc.RegisterNewTask(expectedTask)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedTask.Title, actual.Title)
}

func (suite *TaskUseCaseTestSuite) TestRegisterNewTaskFindAuthor_Fail() {
	suite.aum.On("FindAuthorByID", expectedTask.AuthorId).Return(model.Author{}, fmt.Errorf("error"))
	_, err := suite.tuc.RegisterNewTask(expectedTask)
	assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)
}

func (suite *TaskUseCaseTestSuite) TestRegisterNewTask_EmptyField() {
	suite.aum.On("FindAuthorByID", expectedTask.AuthorId).Return(expectedAuthor, nil)
	payloadMock := model.Task{
		Title:    "This is title",
		Content:  "",
		AuthorId: "1",
	}
	_, err := suite.tuc.RegisterNewTask(payloadMock)
	assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)
}

func (suite *TaskUseCaseTestSuite) TestRegisterNewTask_Fail() {
	suite.aum.On("FindAuthorByID", expectedTask.AuthorId).Return(expectedAuthor, nil)
	suite.trm.On("Create", expectedTask).Return(model.Task{}, fmt.Errorf("error"))
	_, err := suite.tuc.RegisterNewTask(expectedTask)
	assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)
}

// func (suite *TaskUseCaseTestSuite) TestFindTaskByAuthor_Success() {
// 	mockData := []model.Task{expectedTask}

// }

func (suite *TaskUseCaseTestSuite) TestFindAllTask_Success() {
	mockData := []model.Task{expectedTask}
	mockPaging := shared_model.Paging{
		Page:        1,
		RowsPerPage: 1,
		TotalRows:   5,
		TotalPages:  1,
	}
	suite.trm.On("List", 1, 5).Return(mockData, mockPaging, nil)
	actual, paging, err := suite.tuc.FindAllTask(1, 5)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), actual, 1)
	assert.Equal(suite.T(), mockPaging.Page, paging.Page)
}

func TestTaskUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(TaskUseCaseTestSuite))
}
