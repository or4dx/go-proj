package main

import "testing"

func TestSum(t *testing.T) {

	// Sum an array with a len of 5 numbers
	t.Run("collection of 5 numbers", func(t *testing.T) {

		numbers := [5]int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v, got, want, numbers", got, want, numbers)
		}

	})

	// Sum a slice with a dynamic len
	t.Run("collection of any size", func(t *testing.T) {

		numbers := []int{1, 2, 3}

		got := SumDynamic(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}

	})

}
