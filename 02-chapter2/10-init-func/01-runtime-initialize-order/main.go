package main

import "fmt"

var x = xSetter()

func main() {
	/*
		import a library into our code. A library is not
		designed to be executed, it offers data structures, methods, functions, etc.
		Libraries probably do not even have a main package. If this library requires
		some initial configuration before invoked (initialize variables, detect the
		operating system, etc.) it looks impossible.

		init functions that are executed once per package.

		import a package the Go runtime follows this order:
		1. Initialize imported packages recursively.
		2. Initialize and assign values to variables.
		3. Execute init functions.
	*/

	fmt.Println("this is main function")
}

func xSetter() int {
	fmt.Println("xSetter")
	return 42
}

func init() {
	fmt.Println("Init function")
}
