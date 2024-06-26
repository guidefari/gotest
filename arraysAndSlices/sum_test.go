package arraysandslices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		// this errors out because an array's size is part of its type's signature
		// numbers := [5]int{1, 2, 3, 4, 5}
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("collection of 3 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

}

func TestSumAll(t *testing.T) {
	t.Run("one slice", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := SumAll([]int{1, 1, 1})
		want := []int{3}

		// Go does not let you use equality operators with slices.
		// Meaning `if got != want` wouldn't work
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("multiple slice", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSums(t, got, want)
	})

	t.Run("", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 8}, []int{0, 9, 0, 1})
		want := []int{10, 10}
		checkSums(t, got, want)
	})

	t.Run("", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 8}, []int{}, []int{0, 9, 0, 1})
		want := []int{10, 0, 10}
		checkSums(t, got, want)
	})

}
