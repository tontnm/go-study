package main

import "fmt"

func main() {
	result := sum(1, 2)
	fmt.Println("result = ", result)
}

func sum(a, b int) int {
	return a + b
}
