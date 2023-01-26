package main

import "fmt"

func main() {
	/*
			The 01-defer statement pushes a function onto a list. This list of functions is
			executed when the surrounding function ends.
			01-defer is useful to clean up resources allocated to a function.

		FILO
	*/

	defer CloseMsg() // 4

	fmt.Println("Doing something...") // 1

	defer fmt.Println("Certainly closed!!!") // 3

	fmt.Println("Doing something else...") // 2
}

func CloseMsg() {
	fmt.Println("Closed!!!")
}
