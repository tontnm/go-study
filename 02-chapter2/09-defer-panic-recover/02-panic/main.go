package main

import "fmt"

func main() {
	/*
		panic stops the execution flow, executes deferred functions
		and returns control to the calling function.
		A call to panic indicates a situation that
		goes beyond the control of the program.
	*/

	defer fmt.Println("Closed main") // 2

	something()

	fmt.Println("Something was finished") // never show
}

func something() {
	defer fmt.Println("closed something") // 1

	for i := 0; i < 5; i++ {
		fmt.Println(i)

		if i > 2 {
			panic("Panic was called")
		}
	}
}
