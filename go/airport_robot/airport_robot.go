package airportrobot

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.
type Greeter interface {
	LanguageName() string
	Greet(string) string
}

func SayHello(visitorName string, g Greeter) string {
	return "I can speak " + g.LanguageName() + ": " + g.Greet(visitorName)
}

type Italian struct{}

func (i Italian) LanguageName() string {
	return "Italian"
}

func (i Italian) Greet(visitorName string) string {
	return "Ciao " + visitorName + "!"
}

type Portuguese struct{}

func (p Portuguese) LanguageName() string {
	return "Portuguese"
}

func (p Portuguese) Greet(visitorName string) string {
	return "Olá " + visitorName + "!"
}
