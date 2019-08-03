package main

import "fmt"

const (
	englishGreeting = "Hello, "
	spanishGreeting = "Hola, "
	frenchGreeting  = "Bonjour, "
)

func hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	greeting := ""
	switch language {
	case "English":
		greeting = englishGreeting
	case "Spanish":
		greeting = spanishGreeting
	case "French":
		greeting = frenchGreeting
	default:
		greeting = englishGreeting
	}
	return greeting + name
}

func main() {
	fmt.Println(hello("Ryan", "Spanish"))
}
