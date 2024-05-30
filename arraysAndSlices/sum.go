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
	var result []int

	for _, value := range numbersToSum {
		result = append(result, Sum(value))
	}

	return result
}

func SumAllTails(numbersToSum ...[]int) []int {
	var result []int

	for _, value := range numbersToSum {
		if len(value) == 0 {
			result = append(result, 0)
		} else {
			tail := value[1:]
			result = append(result, Sum(tail))
		}

	}

	return result
}
