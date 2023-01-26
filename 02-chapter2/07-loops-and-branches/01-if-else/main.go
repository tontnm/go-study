package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// emulates tossing a coin and tell us if we get head or tail.
	rand.Seed(time.Now().UnixNano())
	x := rand.Float32()

	if x < 0.5 {
		fmt.Println("head")
	} else {
		fmt.Println("tail")
	}

	// no ternary operator in Go
}
