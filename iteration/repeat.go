package iteration

const repeatCount = 5

func Repeat(s string) string {
	var repeated string
	for j := 0; j < repeatCount; j++ {
		repeated += s
	}
	return repeated
}
