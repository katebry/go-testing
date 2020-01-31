package hello

const greeting string = "hello"

// World returns simple string
func World(name string) string {
	return greeting + name
}
