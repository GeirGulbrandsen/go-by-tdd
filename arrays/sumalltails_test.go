package main

import (
	"reflect"
	"testing"
)

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d", got, want)
		}
	}
	t.Run("can sum the tail of one slice", func(t *testing.T) {
		got := SumAllTails([]int{1, 2})
		want := []int{2}

		checkSums(t, got, want)

	})
	t.Run("can sum the tail of three slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{3, 4}, []int{5, 6})
		want := []int{2, 4, 6}

		checkSums(t, got, want)

	})
	t.Run("can handle empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5, 6})
		want := []int{0, 15}

		checkSums(t, got, want)
	})
}
