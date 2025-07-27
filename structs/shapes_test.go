package structs

import "testing"

func TestPerimeter(t *testing.T) {
	r := Rectangle{Width: 4.5, Height: 5.5}
	got := Perimeter(r)
	want := 20.0
	if got != want {
		t.Errorf("got = %v; want %v", got, want)
	}
}

func TestArea(t *testing.T) {
	r := Rectangle{Width: 4, Height: 5}
	got := Area(r)
	want := 20.0
	if got != want {
		t.Errorf("got = %v; want %v", got, want)
	}
}
