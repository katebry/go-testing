package hello

import "testing"

// TestWorld tests that a string is returned with a name based on input
func TestWorld(t *testing.T) {

	t.Run("Say hello 'your name!' when a name is supplied", func(t *testing.T) {

		got := World(" kate!")
		want := "hello kate!"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("Say hello world! when an empty string is supplied", func(t *testing.T) {

		got := World("")
		want := "hello world!"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
