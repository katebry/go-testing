package greet

import "fmt"

const englishGreeting string = "hello"
const spanishGreeting string = "hola"
const frenchGreeting string = "bonjour"
const germanGreeting string = "hallo"

const spanish = "Spanish"
const french = "French"
const german = "German"

// Greeter returns a greeting in the relevant language
func Greeter(name string, language string) string {
	if name == "" {
		name = " world"
	}
	greeting := greetingPrefix(language) + " " + name
	fmt.Println(greeting)
	return greeting
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishGreeting
	case french:
		prefix = frenchGreeting
	case german:
		prefix = germanGreeting
	default:
		prefix = englishGreeting
	}
	return
}
