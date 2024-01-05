package usecase

import (
	"gotodo-app/mock/service_mock"
	"gotodo-app/mock/usecase_mock"
	"gotodo-app/model"
	"gotodo-app/model/dto"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type AuthUseCaseTestSuite struct {
	suite.Suite
	aum *usecase_mock.AuthorUseCaseMock
	jsm *service_mock.JwtServiceMock
	au  AuthUseCase
}

func (suite *AuthUseCaseTestSuite) SetupTest() {
	suite.aum = new(usecase_mock.AuthorUseCaseMock)
	suite.jsm = new(service_mock.JwtServiceMock)
	suite.au = NewAuthUseCase(suite.aum, suite.jsm)
}

func (suite *AuthUseCaseTestSuite) TestLogin_Success() {
	mockLogin := dto.AuthRequestDto{
		Email:    "mail@mail.com",
		Password: "password",
	}
	mockAuthor := model.Author{
		ID:       "1",
		Email:    "mail@mail.com",
		Password: "password",
		Role:     "user",
	}
	mockAuthResponse := dto.AuthResponseDto{
		Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",
	}
	suite.aum.On("FindAuthorByEmail", mockLogin.Email).Return(mockAuthor, nil)
	suite.jsm.On("CreateToken", mockAuthor).Return(mockAuthResponse, nil)
	actual, err := suite.au.Login(mockLogin)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), mockAuthResponse, actual)
}

func TestAuthUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(AuthUseCaseTestSuite))
}
