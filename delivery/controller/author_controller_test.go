package controller

import (
	"fmt"
	"gotodo-app/mock/middleware_mock"
	"gotodo-app/mock/usecase_mock"
	"gotodo-app/model"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type AuthorControllerTestSuite struct {
	suite.Suite
	rg  *gin.RouterGroup
	aum *usecase_mock.AuthorUseCaseMock
	amm *middleware_mock.AuthMiddlewareMock
}

func (suite *AuthorControllerTestSuite) SetupTest() {
	suite.aum = new(usecase_mock.AuthorUseCaseMock)
	suite.amm = new(middleware_mock.AuthMiddlewareMock)
	r := gin.Default()
	gin.SetMode(gin.TestMode)
	rg := r.Group("/api/v1")
	rg.Use(suite.amm.RequireToken("userssss"))
	suite.rg = rg
}

// func (suite *AuthorControllerTestSuite) TestGetHandler_Success() {
// 	mockAuthor := []model.Author{
// 		{
// 			ID:        "1",
// 			Name:      "Name",
// 			Email:     "mail@mail.com",
// 			CreatedAt: time.Now(),
// 			UpdatedAt: time.Now(),
// 			Tasks: []model.Task{
// 				{
// 					ID:        "1",
// 					Title:     "Title",
// 					Content:   "Content",
// 					CreatedAt: time.Now(),
// 					UpdatedAt: time.Now(),
// 				},
// 			},
// 		},
// 	}
// 	suite.aum.On("FindAuthorByID", mockAuthor[0].ID).Return(mockAuthor[0], nil)
// 	authorController := NewAuthorController(suite.aum, suite.rg, suite.amm)
// 	authorController.Route()
// 	request, err := http.NewRequest(http.MethodGet, "/api/v1/authors/1", nil)
// 	assert.NoError(suite.T(), err)
// 	record := httptest.NewRecorder()
// 	ctx, _ := gin.CreateTestContext(record)
// 	ctx.Request = request
// 	ctx.Set("author", mockAuthor[0].ID)
// 	authorController.getHandler(ctx)
// 	assert.Equal(suite.T(), http.StatusOK, record.Code)
// }

// func (suite *AuthorControllerTestSuite) TestGetHandler_Fail() {
// 	mockAuthor := []model.Author{
// 		{
// 			ID:        "1",
// 			Name:      "Name",
// 			Email:     "mail@gmail.com",
// 			CreatedAt: time.Now(),
// 			UpdatedAt: time.Now(),
// 			Tasks: []model.Task{
// 				{
// 					ID:        "1",
// 					Title:     "Title",
// 					Content:   "Content",
// 					CreatedAt: time.Now(),
// 					UpdatedAt: time.Now(),
// 				},
// 			},
// 		},
// 	}
// 	suite.aum.On("FindAuthorByID", mockAuthor[0].ID).Return(mockAuthor[0], fmt.Errorf("error"))
// 	authorController := NewAuthorController(suite.aum, suite.rg, suite.amm)
// 	authorController.Route()
// 	request, err := http.NewRequest(http.MethodGet, "/api/v1/authors/1", nil)
// 	assert.NoError(suite.T(), err)
// 	record := httptest.NewRecorder()
// 	ctx, _ := gin.CreateTestContext(record)
// 	ctx.Request = request
// 	ctx.Set("author", mockAuthor[0].ID)
// 	authorController.getHandler(ctx)
// 	assert.Equal(suite.T(), http.StatusInternalServerError, record.Code)
// }

func (suite *AuthorControllerTestSuite) TestListHandler_Success() {
	mockAuthor := []model.Author{
		{
			ID:        "1",
			Name:      "Name",
			Email:     "mail@mail.com",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Tasks: []model.Task{
				{
					ID:        "1",
					Title:     "Title",
					Content:   "Content",
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
			},
		},
	}
	suite.aum.On("FindAllAuthor", mockAuthor[0].ID).Return(mockAuthor, nil)
	authorController := NewAuthorController(suite.aum, suite.rg, suite.amm)
	authorController.Route()
	request, err := http.NewRequest(http.MethodGet, "/api/v1/authors", nil)
	assert.NoError(suite.T(), err)
	record := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = request
	ctx.Set("author", mockAuthor[0].ID)
	authorController.listHandler(ctx)
	assert.Equal(suite.T(), http.StatusOK, record.Code)
}

func (suite *AuthorControllerTestSuite) TestListHandler_Fail() {
	mockTokenJwt := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"

	suite.aum.On("FindAllAuthor", "xx").Return([]model.Author{}, fmt.Errorf("error"))
	authorController := NewAuthorController(suite.aum, suite.rg, suite.amm)
	authorController.Route()
	record := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodGet, "/api/v1/authors", nil)
	assert.NoError(suite.T(), err)
	request.Header.Set("Authorization", "Bearer "+mockTokenJwt)
	ctx, _ := gin.CreateTestContext(record)
	ctx.Request = request
	ctx.Set("author", "xx")
	authorController.listHandler(ctx)
	assert.Equal(suite.T(), http.StatusInternalServerError, record.Code)
}

func TestAuthorControllerTestSuite(t *testing.T) {
	suite.Run(t, new(AuthorControllerTestSuite))
}
