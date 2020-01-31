package hello

const greeting string = "hello"

// World returns simple string
func World(name string) string {
	if len(name) > 0 {
		return greeting + name
	}
	return greeting + " world!"
}
