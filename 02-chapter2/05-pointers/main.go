package main

import "fmt"

func main() {
	x := 100

	a(x)
	fmt.Println(x)

	b(&x)
	fmt.Println(x)

	fmt.Println(&x)
}

// pass value
func a(i int) { // i is a copy
	i = 0
}

// pass reference
func b(i *int) {
	*i = 0
}
