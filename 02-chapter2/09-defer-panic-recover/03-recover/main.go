package main

import "fmt"

func main() {
	defer fmt.Println("closed main") // 2

	something()

	fmt.Println("main finished") // 3
}

func something() {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("No need to panic if i =", r) // 1
		}
	}()

	for i := 0; i < 5; i++ {
		fmt.Println(i)

		if i > 2 {
			panic(i) // giá trị truyền vào panic là r ở trên
		}
	}

	fmt.Println("Closed something normally")
}
