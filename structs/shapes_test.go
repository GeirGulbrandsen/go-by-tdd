package structs

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(4.5, 5.5)
	want := 20.0
	if got != want {
		t.Errorf("got = %v; want %v", got, want)
	}
}

func TestArea(t *testing.T) {
	got := Area(4, 5)
	want := 20.0
	if got != want {
		t.Errorf("got = %v; want %v", got, want)
	}
}
