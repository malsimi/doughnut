package integers

// Add takes two integers and returns the sum of them.
func Add(a, b int) int {
	return a + b
}

func Sum(numbers []int) int {
	sum := 0
	for _, value := range numbers {
		sum += value
	}
	return sum
}

func SumAll(numbers ...[]int) (sums []int) {
	sums = make([]int, len(numbers))

	for i := range numbers {
		sums[i] = Sum(numbers[i])
	}

	return

}

func SumAllTails(numbers ...[]int) (sums []int) {

	sums = make([]int, len(numbers))

	for i := range numbers {
		tail := numbers[i][1:]
		sums[i] = Sum(tail)
	}

	return
}
