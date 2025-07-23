package main

func SumAllTails(i ...[]int) []int {
	var tails []int
	for _, v := range i {
		if len(v) == 0 {
			tails = append(tails, 0)
			continue
		}
		tails = append(tails, Sum(v[1:]))
	}
	return tails
}
