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

	AreaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10.0}, 314.1592653589793},
	}

	for _, tt := range AreaTests {
		t.Run("calculate area", func(t *testing.T) {
			checkArea(t, tt.shape, tt.want)
		})
	}
}
