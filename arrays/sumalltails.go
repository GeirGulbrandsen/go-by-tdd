package main

func SumAllTails(i ...[]int) []int {
	var tails []int
	for _, v := range i {
		tails = append(tails, Sum(v[1:]))
	}
	return tails
}
