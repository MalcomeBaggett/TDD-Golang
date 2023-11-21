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