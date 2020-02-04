package hello

const englishGreeting string = "hello"
const spanishGreeting string = "hola"
const frenchGreeting string = "bonjour"

const spanish = "Spanish"
const french = "French"

// World returns simple string
func World(name string, language string) string {
	if name == "" {
		name = " world"
	}
	if language == spanish {
		return spanishGreeting + name
	}
	if language == french {
		return frenchGreeting + name
	}

	return englishGreeting + name
}
