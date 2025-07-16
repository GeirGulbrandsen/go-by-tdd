package iteration

import "strings"

const repeatCount = 5

func Repeat(s string) string {
	var repeated strings.Builder
	for j := 0; j < repeatCount; j++ {
		repeated.WriteString(s)
	}
	return repeated.String()
}
