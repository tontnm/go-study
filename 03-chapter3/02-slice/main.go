package main

import (
	"fmt"
	"reflect"
)

func main() {
	// slice is a reference to an array,
	/*
		a slice can reserve a larger amount of memory that
		does not necessarily have to be filled. The filled
		memory corresponds to its length, and all the available
		memory is the capacity.
	*/

	a := [5]string{"a", "b", "c", "d", "e"}

	fmt.Println(a)          // [a,b,c,d,e]
	fmt.Println(a[:])       // [a,b,c,d,e]
	fmt.Println(a[0])       // a
	fmt.Println(a[0], a[1]) // a, b
	fmt.Println(a[0:2])     // [a,b]
	fmt.Println(a[1:4])     // [b,c,d]
	fmt.Println(a[:2])      // [a,b]
	fmt.Println(a[2:])      // [c,d,e]

	fmt.Println(reflect.TypeOf(a))
	fmt.Println(reflect.TypeOf(a[0:3]))
	fmt.Println(reflect.TypeOf(a[0]))

	b := []int{0, 1, 2, 3, 4}
	fmt.Println(b, len(b), cap(b)) // [0,1,2,3,4] 5 5

	c := append(b, 5)
	fmt.Println(c, len(c), cap(c)) // [0,1,2,3,4,5] 6 10

	c = append(c, 6)
	fmt.Println(c, len(c), cap(c)) // [0,1,2,3,4,5,6] 7 10

	d := c[1:4]
	fmt.Println(d, len(d), cap(d)) // [1,2,3] 3 9

	e := make([]int, 5, 10)
	fmt.Println(e, len(e), cap(e)) // [0,0,0,0,0] 5 10

	// e[6]=5 will fail

	names := []string{"a", "b", "c"}

	for i := 0; i < len(names); i++ {
		fmt.Println(i, names[i])
	}

	for position, name := range names {
		fmt.Println(position, name)
	}

	for _, name := range names {
		// name is a copy
		name = name + "_changed"
	}

	fmt.Println(names)

	for position, name := range names {
		// this modifies the original value
		names[position] = name + "_changed"
	}
	
	fmt.Println(names)
}
