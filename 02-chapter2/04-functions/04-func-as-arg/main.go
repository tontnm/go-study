package main

import "fmt"

func main() {
	result1 := calculate(divide, 6, 2)
	result2 := calculate(multiply, 6, 2)

	fmt.Println(result1)
	fmt.Println(result2)
}

func calculate(operator func(int, int) int, a, b int) int {
	return operator(a, b)
}

func divide(a, b int) int {
	return a / b
}

func multiply(a, b int) int {
	return a * b
}
