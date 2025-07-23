package main

func SumAll(slicesToSum ...[]int) []int {
	var results []int

	for _, numbers := range slicesToSum {
		results = append(results, Sum(numbers))
	}

	return results
}
