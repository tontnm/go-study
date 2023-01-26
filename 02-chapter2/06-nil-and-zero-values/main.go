package main

import "fmt"

func main() {
	// Zero values during initialization.

	var a int
	fmt.Println(a) // 0

	var b string
	fmt.Println(b) // ""

	var c bool
	fmt.Println(c) // false

	var d *int
	fmt.Println(d) // nil

	var e func()
	fmt.Println(e) // nil
}
