package structs

import "testing"

func TestArea(t *testing.T) {
	got := Area(4, 5)
	want := 20.0
	if got != want {
		t.Errorf("got = %v; want %v", got, want)
	}
}

func Area(f1, f2 float64) float64 {
	return f1 * f2
}
