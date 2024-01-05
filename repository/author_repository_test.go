package repository

import (
	"database/sql"
	"fmt"
	"gotodo-app/model"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var expectedAuthor = model.Author{
	ID:        "1",
	Name:      "Jack",
	Email:     "jack@mail.com",
	Password:  "password",
	Role:      "admin",
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
	Tasks:     nil,
}

var expectedTask = model.Task{
	ID:        "1",
	Title:     "This is title",
	Content:   "This is content",
	AuthorId:  expectedAuthor.ID,
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
}

type AuthorRepositoryTestSuite struct {
	suite.Suite
	mockDb  *sql.DB
	mockSql sqlmock.Sqlmock
	repo    AuthorRepository
}

func (suite *AuthorRepositoryTestSuite) SetupTest() {
	db, mock, _ := sqlmock.New()
	suite.mockDb = db
	suite.mockSql = mock
	suite.repo = NewAuthorRepository(suite.mockDb)
}

func (suite *AuthorRepositoryTestSuite) TestGetByEmail_Success() {
	rows := sqlmock.NewRows([]string{"id", "email", "password", "role"}).
		AddRow(expectedAuthor.ID, expectedAuthor.Email, expectedAuthor.Password, expectedAuthor.Role)
	suite.mockSql.ExpectQuery(regexp.QuoteMeta(`SELECT id, email, password, role FROM authors WHERE email = $1`)).WithArgs(expectedAuthor.Email).WillReturnRows(rows)
	actualAuthor, actualError := suite.repo.GetByEmail(expectedAuthor.Email)
	assert.Nil(suite.T(), actualError)
	assert.NoError(suite.T(), actualError)
	assert.Equal(suite.T(), expectedAuthor.Email, actualAuthor.Email)
}

func (suite *AuthorRepositoryTestSuite) TestGetByEmail_Fail() {
	suite.mockSql.ExpectQuery(regexp.QuoteMeta(`SELECT id, email, password, role FROM authors WHERE email = $1`)).WillReturnError(fmt.Errorf("error"))
	actualAuthor, actualError := suite.repo.GetByEmail("xx@xx.com")
	assert.Error(suite.T(), actualError)
	assert.NotNil(suite.T(), actualError)
	assert.Equal(suite.T(), model.Author{}, actualAuthor)
}

func (suite *AuthorRepositoryTestSuite) TestGet_Success() {
	suite.mockSql.ExpectQuery(`SELECT`).WithArgs(expectedAuthor.ID).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "role", "created_at", "updated_at"}).AddRow(expectedAuthor.ID, expectedAuthor.Name, expectedAuthor.Email, expectedAuthor.Role, expectedAuthor.CreatedAt, expectedAuthor.UpdatedAt))

	suite.mockSql.ExpectQuery(`SELECT`).WithArgs(expectedAuthor.ID).WillReturnRows(sqlmock.NewRows([]string{"id", "title", "content", "created_at", "updated_at"}).AddRow(expectedTask.ID, expectedTask.Title, expectedTask.Content, expectedTask.CreatedAt, expectedTask.UpdatedAt))

	// simulate append task
	expectedAuthor.Tasks = append(expectedAuthor.Tasks, expectedTask)

	actual, err := suite.repo.Get(expectedAuthor.ID)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedAuthor.Name, actual.Name)
	assert.Len(suite.T(), expectedAuthor.Tasks, 1)
}

func (suite *AuthorRepositoryTestSuite) TestGet_QueryRowAuthorFail() {
	// Simulate query row author error
	suite.mockSql.ExpectQuery(regexp.QuoteMeta(`SELECT`)).WithArgs(expectedAuthor.ID).WillReturnError(fmt.Errorf("error"))
	_, err := suite.repo.Get(expectedAuthor.ID)
	assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)
}

func (suite *AuthorRepositoryTestSuite) TestGet_QueryTaskFail() {
	suite.mockSql.ExpectQuery(regexp.QuoteMeta(`SELECT`)).WithArgs(expectedAuthor.ID).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "role", "created_at", "updated_at"}).AddRow(expectedAuthor.ID, expectedAuthor.Name, expectedAuthor.Email, expectedAuthor.Role, expectedAuthor.CreatedAt, expectedAuthor.UpdatedAt))

	suite.mockSql.ExpectQuery(regexp.QuoteMeta(`SELECT`)).WithArgs(expectedAuthor.ID).WillReturnError(fmt.Errorf("error"))
	_, err := suite.repo.Get(expectedAuthor.ID)
	assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)
}

func (suite *AuthorRepositoryTestSuite) TestGet_ScanFail() {
	suite.mockSql.ExpectQuery(regexp.QuoteMeta(`SELECT`)).WithArgs(expectedAuthor.ID).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "email", "role", "created_at", "updated_at"}).AddRow(expectedAuthor.ID, expectedAuthor.Name, expectedAuthor.Email, expectedAuthor.Role, expectedAuthor.CreatedAt, expectedAuthor.UpdatedAt))

	suite.mockSql.ExpectQuery(regexp.QuoteMeta(`SELECT t.id, t.title, t.content, t.created_at, t.updated_at FROM authors a JOIN tasks t on a.id = t.author_id
WHERE a.id = $1`)).WithArgs(expectedAuthor.ID).WillReturnRows(sqlmock.NewRows([]string{"id", "title"}).AddRow(expectedTask.ID, expectedTask.Title))
	_, err := suite.repo.Get(expectedAuthor.ID)
	assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)

}

func TestAuthorRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(AuthorRepositoryTestSuite))
}
