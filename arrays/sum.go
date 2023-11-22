package arrays

import "fmt"

func Sum(numbers []int) int {
	sum := 0
	for index, number := range numbers {
		sum += number
		fmt.Println(index)
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)

	sums := make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}
	return sums
}