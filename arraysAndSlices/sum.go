package arraysandslices

func Sum(numbers []int) int {

	sum := 0

	// range lets you iterate over an array. On each iteration, range returns two values - the index and the value.
	// We are choosing to ignore the index value by using _ blank identifier.
	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	result := make([]int, lengthOfNumbers)

	for index, value := range numbersToSum {
		result[index] = Sum(value)
	}

	return result
}
