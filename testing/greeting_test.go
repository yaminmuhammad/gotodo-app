package testing

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type GreetingServiceTestSuite struct {
	suite.Suite
	service     GreetingService
	mockService *GreetingServiceMock
}

func (suite *GreetingServiceTestSuite) SetupTest() {
	suite.mockService = new(GreetingServiceMock)
	suite.service = suite.mockService // gunakan ini jika tidak ingin memperdulikan fungsionalitas
	// suite.service = NewGreetingService() // gunakan ini jika ingin melakukan simulasi bagaimana fungsi asli ini diimplementasikan seperti real case
}

func (suite *GreetingServiceTestSuite) TestSayHello_Success() {
	mockPerson := Person{Name: "Jack"}
	suite.mockService.On("SayHello", mockPerson).Return(mockPerson, nil)
	actual, err := suite.service.SayHello(mockPerson)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), mockPerson, actual)
}

func (suite *GreetingServiceTestSuite) TestSayHello_EmptyName() {
	mockPerson := Person{Name: ""}
	suite.mockService.On("SayHello", mockPerson).Return(Person{}, fmt.Errorf("error"))
	actual, err := suite.service.SayHello(mockPerson)
	assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), mockPerson, actual)
}

func TestGreetingServiceTestSuite(t *testing.T) {
	suite.Run(t, new(GreetingServiceTestSuite))
}
