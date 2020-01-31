package hello

import "testing"

// TestWorldKate tests that a string is returned with a name based on input
func TestWorld(t *testing.T) {
	got := World(" kate!")
	want := "hello kate!"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
