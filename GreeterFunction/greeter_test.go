package greet

import "testing"

// TestGreeter tests that a greeting is returned in the correct language based on input
func TestGreeter(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("Say hello 'your name!' when a name is supplied", func(t *testing.T) {
		got := Greeter(" kate", "")
		want := "hello kate"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Say hello world! when an empty string is supplied", func(t *testing.T) {

		got := Greeter("", "")
		want := "hello world"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Say hello in Spanish", func(t *testing.T) {

		got := Greeter(" kate", "Spanish")
		want := "hola kate"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Say hello in French", func(t *testing.T) {

		got := Greeter(" kate", "French")
		want := "bonjour kate"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Say hello in German", func(t *testing.T) {

		got := Greeter(" kate", "German")
		want := "hallo kate"
		assertCorrectMessage(t, got, want)
	})
}
