package structs

type Rectangle struct {
	Width  float64
	Height float64
}

func Perimeter(r Rectangle) float64 {
	return 2 * (r.Height + r.Width)
}

func Area(f1, f2 float64) float64 {
	return f1 * f2
}
