package structs

import "testing"

func TestPerimeter(t *testing.T) {
	r := Rectangle{4.5, 5.5}
	got := Perimeter(r)
	want := 20.0
	if got != want {
		t.Errorf("got = %v; want %v", got, want)
	}
}

func TestArea(t *testing.T) {
	t.Run("calculate area of rectangle", func(t *testing.T) {
		r := Rectangle{4.0, 5.5}
		got := Area(r)
		want := 22.0
		if got != want {
			t.Errorf("got = %v; want %v", got, want)
		}
	})
}
