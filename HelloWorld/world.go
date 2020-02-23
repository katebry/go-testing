package hello

const englishGreeting string = "hello"
const spanishGreeting string = "hola"
const frenchGreeting string = "bonjour"
const germanGreeting string = "hallo"

const spanish = "Spanish"
const french = "French"
const german = "German"

// World returns simple string
func World(name string, language string) string {
	if name == "" {
		name = " world"
	}
	return greetingPrefix(language) + name
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
