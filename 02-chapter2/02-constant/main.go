package main

import (
	"fmt"
	"reflect"
)

const (
	Pi               = 3.14 // default Go will select the largest available type
	Avogadro float32 = 6.022e23
)

func main() {
	fmt.Println("Pi", Pi)
	fmt.Println("Pi type", reflect.TypeOf(Pi)) // float64

	fmt.Println("Avogadro", Avogadro)
	fmt.Println("Avogadro type", reflect.TypeOf(Avogadro)) // float32

}
