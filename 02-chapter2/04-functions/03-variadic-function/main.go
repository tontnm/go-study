package main

import "fmt"

func main() {
	result1 := sum(0, 1, 2, 3, 4, 5)

	nums := []int{0, 1, 2, 3, 4, 5}

	result2 := sum(nums...)

	fmt.Printf("result1=%d\n", result1)
	fmt.Printf("result2=%d", result2)
}

func sum(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total = total + number
	}

	return total
}
