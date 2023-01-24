package main

import (
	"fmt"
	"os"
)

func main() {
	argsWithProgram := os.Args
	argsWithoutProgram := os.Args[1:]

	arg := os.Args[3]

	fmt.Println(argsWithProgram)    // [main hello world tony]
	fmt.Println(argsWithoutProgram) // [hello world tony]
	fmt.Println(arg)                // tony
}

// go run main.go hello world tony
