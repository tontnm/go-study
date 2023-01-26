package main

import "fmt"

func main() {
	sum, sub := ops(2, 2)
	fmt.Printf("2+2=%d ; 2-2=%d\n", sum, sub)

	b, _ := ops(10, 2)
	fmt.Printf("10+2=%d", b)
}

func ops(a, b int) (int, int) {
	return a + b, a - b
}
