package arraysandslices

import "testing"

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
		// this errors out because an array's size is part of its type's signature
		// numbers := [5]int{1, 2, 3, 4, 5}
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
		// this errors out because an array's size is part of its type's signature
		// numbers := [5]int{1, 2, 3, 4, 5}
		numbers := []int{1, 2, 3}

		got := SumAll([]int{1, 1, 1})
		want := []int{3}

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("multiple slice", func(t *testing.T) {
		// this errors out because an array's size is part of its type's signature
		// numbers := [5]int{1, 2, 3, 4, 5}
		numbers := []int{1, 2, 3}

		got := SumAll([]int{1, 2}, []int{0, 9})
		want := []int{3, 9}

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}
