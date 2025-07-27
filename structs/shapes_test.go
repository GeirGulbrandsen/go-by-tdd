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

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got = %g; want %g", got, want)
		}
	}

	t.Run("calculate area of rectangle", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		checkArea(t, rectangle, 72.0)
	})
	t.Run("calculate area of circle", func(t *testing.T) {
		circle := Circle{10.0}
		checkArea(t, circle, 314.1592653589793)
	})
}
