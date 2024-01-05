package repository

import (
	"database/sql"
	"fmt"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TaskRepositoryTestSuite struct {
	suite.Suite
	mockDb  *sql.DB
	mockSql sqlmock.Sqlmock
	repo    TaskRepository
}

func (suite *TaskRepositoryTestSuite) SetupTest() {
	db, mock, _ := sqlmock.New()
	suite.mockDb = db
	suite.mockSql = mock
	suite.repo = NewTaskRepository(suite.mockDb)
}

func (suite *TaskRepositoryTestSuite) TestCreate_Success() {
	suite.mockSql.ExpectQuery(`INSERT INTO tasks`).WithArgs(
		expectedTask.Title,
		expectedTask.Content,
		expectedTask.AuthorId).WillReturnRows(sqlmock.NewRows([]string{"id", "created_at"}).AddRow(expectedTask.ID, expectedTask.CreatedAt))
	actual, err := suite.repo.Create(expectedTask)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectedTask.Title, actual.Title)
}

func (suite *TaskRepositoryTestSuite) TestCreate_Fail() {
	suite.mockSql.ExpectQuery(`INSERT INTO tasks`).WithArgs(expectedTask.Title, expectedTask.Content, expectedTask.AuthorId, expectedTask.UpdatedAt).WillReturnError(fmt.Errorf("error"))
	_, err := suite.repo.Create(expectedTask)
	assert.Error(suite.T(), err)
	assert.NotNil(suite.T(), err)
}

func (suite *TaskRepositoryTestSuite) TestGetByAuthor_Success() {
	rows := sqlmock.NewRows([]string{"id", "title", "content", "author_id", "created_at", "updated_at"}).
		AddRow(expectedTask.ID, expectedTask.Title, expectedTask.Content, expectedTask.AuthorId, expectedTask.CreatedAt, expectedTask.UpdatedAt)
	suite.mockSql.ExpectQuery(regexp.QuoteMeta(`SELECT id, title, content, author_id, created_at, updated_at FROM tasks WHERE author_id = $1`)).WithArgs(expectedTask.AuthorId).WillReturnRows(rows)
	_, err := suite.repo.GetByAuthor(expectedTask.AuthorId)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
}

func (suite *TaskRepositoryTestSuite) TestGetByAuthor_Fail() {
	suite.mockSql.ExpectQuery(regexp.QuoteMeta(`SELECT id, title, content, author_id, created_at, updated_at FROM tasks WHERE  author_id = $1`)).WillReturnError(fmt.Errorf("error"))
	_, err := suite.repo.GetByAuthor("xxx")
	assert.Error(suite.T(), err)
	assert.NotNil(suite.T(), err)
}

func (suite *TaskRepositoryTestSuite) TestGetByAuthorRow_Fail() {
	suite.mockSql.ExpectQuery(regexp.QuoteMeta(`SELECT id, title, content, author_id, created_at, updated_at FROM tasks WHERE author_id = $1`)).WithArgs("1").WillReturnRows(sqlmock.NewRows([]string{"id", "title"}).AddRow(expectedTask.ID, expectedTask.Title))

	_, err := suite.repo.GetByAuthor("1")
	assert.Error(suite.T(), err)
	assert.NotNil(suite.T(), err)
}

func TestTaskRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(TaskRepositoryTestSuite))
}
