package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// take action for several values of the same variable.
	// takes a number and prints the name of the corresponding finger.

	var finger int = 1

	switch finger {
	case 0:
		fmt.Println("Thumb")
	case 1:
		fmt.Println("Index")
	case 2:
		fmt.Println("Middle")
	case 3:
		fmt.Println("Ring")
	case 4:
		fmt.Println("Pinkie")
	default:
		fmt.Println("Humans only have 5 fingers")
	}

	rand.Seed(time.Now().UnixNano())
	x := rand.Float32()

	switch {
	case x < 0.25:
		fmt.Println("Q1")
	case x < 0.5:
		fmt.Println("Q2")
	case x < 0.75:
		fmt.Println("Q3")
	default:
		fmt.Println("Q4")
	}

	/* When several cases share the same logic, they can be stacked.
	switch {
	case 0:
	case 1:
		// statement for 0 and 1
	default:
		// default statement

	Using default in every switch is a good practice to avoid errors and unexpected
	behaviours.
	*/

}
