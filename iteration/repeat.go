package iteration

import "strings"

func Repeat(s string, i int) string {
	var repeated strings.Builder
	for range i {
		repeated.WriteString(s)
	}
	return repeated.String()
}
