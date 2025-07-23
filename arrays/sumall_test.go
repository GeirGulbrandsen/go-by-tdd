package main

import (
	"reflect"
	"testing"
)

func TestSumAll(t *testing.T) {
	t.Run("can sum one slice", func(t *testing.T) {
		got := SumAll([]int{1, 2})
		want := []int{3}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d", got, want)
		}
	})
	t.Run("can sum three slicea", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{3, 4}, []int{5, 6})
		want := []int{3, 7, 11}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d", got, want)
		}
	})

}
