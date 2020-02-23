package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q\n", expected, repeated)
	}
}
