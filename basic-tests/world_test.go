package hello

import "testing"

// TestWorld tests that a string is returned with a name based on input
func TestWorld(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("Say hello 'your name!' when a name is supplied", func(t *testing.T) {
		got := World(" kate", "")
		want := "hello kate"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Say hello world! when an empty string is supplied", func(t *testing.T) {

		got := World("", "")
		want := "hello world"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Say hello in Spanish", func(t *testing.T) {

		got := World(" kate", "Spanish")
		want := "hola kate"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Say hello in French", func(t *testing.T) {

		got := World(" kate", "French")
		want := "bonjour kate"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Say hello in German", func(t *testing.T) {

		got := World(" kate", "German")
		want := "hallo kate"
		assertCorrectMessage(t, got, want)
	})
}
