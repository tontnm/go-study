package main

import "fmt"

func main() {
	// map creation.
	var ages map[string]int
	fmt.Println(ages)

	// this fails, ages was not initialized
	// ages["a"]=33

	ages = make(map[string]int, 5)
	// now it works because it was initialized
	ages["a"] = 33

	fmt.Println(ages)

	names := make(map[string]int, 5)
	names["a"] = 22

	fmt.Println(names)

	/*
		map can be retrieved using any key type value this operation returns two items,
		one with the expected value and true if the item was found. In case, the key
		was not found, the value would be nil.
	*/
	fmt.Println(names, len(names)) // 1

	name1, found1 := names["a"]
	fmt.Println(name1, found1) // 22,true

	name2, found2 := names["b"]
	fmt.Println(name2, found2) // 0, false

	// add item to map
	names["b"] = 11
	fmt.Println(names, len(names)) // 2

	// update
	names["a"] = 2
	fmt.Println(names, len(names)) // 2

	// delete
	delete(names, "b")
	fmt.Println(names, len(names)) // 1

	//The function range does not guarantee the same order for consecutive iterations.
	for k, v := range names {
		fmt.Printf("%s\t\t%d\n", k, v)
	}
}
