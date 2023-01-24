package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	argsWithProgram := os.Args

	num1, err := strconv.Atoi(argsWithProgram[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(2) // zero indicates success, non-zero an error, range [0, 125].
	}

	num2, err := strconv.Atoi(argsWithProgram[2])
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	result := num1 + num2

	fmt.Printf("%d + %d = %d", num1, num2, result)
}
