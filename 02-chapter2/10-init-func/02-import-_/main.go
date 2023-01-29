package main

import (
	_ "02-import-_/a" // 2, 3
	"fmt"             // 1
)

func init() { // 4
	fmt.Println("Init from my program")
}

func main() { // 5
	fmt.Println("main program")
}
