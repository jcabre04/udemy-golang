package main

import "fmt"

type bot interface {
	// Any other type with a function with this name and return type
	// is now an honorary "bot" type'

	//a method set
	getGreeting() string
}

type englishBot struct {}
type spanishBot struct {}

func (englishBot) getGreeting() string {
	// VERY custom logic for generating an english greeting
	return "Hello!"
}

func (spanishBot) getGreeting() string {
	// VERY custom logic for generating a spanish greeting
	return "Hola!"
}

// All bot types can use this function
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}