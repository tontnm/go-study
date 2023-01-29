package main

import "fmt"

func main() {
	// arrays are typed and their size is fixed.
	// an array allocates a fixed amount of memory that is directly related to its length,

	var a [5]int
	fmt.Println(a)

	b := [5]int{0, 1, 2, 3, 4}
	fmt.Println(b)

	c := [5]int{0, 1, 2}
	fmt.Println(c)
}
