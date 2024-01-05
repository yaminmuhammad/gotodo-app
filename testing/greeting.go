package testing

import "fmt"

type Person struct {
	Name string
}

type GreetingService interface {
	SayHello(person Person) (Person, error)
}

type greetingService struct{}

func (g *greetingService) SayHello(person Person) (Person, error) {
	if person.Name == "" {
		return Person{}, fmt.Errorf("name can't be empty")
	}
	return Person{Name: person.Name}, nil
}

func NewGreetingService() GreetingService { return &greetingService{} }
