package testing

import (
	"github.com/stretchr/testify/mock"
)

type GreetingServiceMock struct {
	mock.Mock
}

func (g *GreetingServiceMock) SayHello(person Person) (Person, error) {
	args := g.Called(person)
	return args.Get(0).(Person), args.Error(1)
}
