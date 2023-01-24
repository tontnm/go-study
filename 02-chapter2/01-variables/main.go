package main

import "fmt"

func main() {
	var a int
	a = 42

	fmt.Println("a: ", a) // 42, nếu ko có dòng này a ở line 6 sẽ báo lỗi vì chưa đc sử dụng

	var aa int = 100

	fmt.Println("aa: ", aa)

	b := -42

	c := "this is a string"

	fmt.Println("b and c: ", b, c)

	var d, e string
	d, e = "var d", "bar e"

	f, g := true, false

	fmt.Println("d,e,f,g: ", d, e, f, g)

	var aRune rune = '€'

	fmt.Println(aRune)        // 8364
	fmt.Printf("%U\n", aRune) // U+20AC
	fmt.Printf("%c\n", aRune) // €

}

/*
	Type Description
		bool Boolean
		string String of characters
		int, int8, int16, int32, int64 Signed integers
		uint, uint8, uint16, uint32, uint64 Unsigned integers
		byte Byte, similar to uint8
		rune(int32) Unicode code point
		float32, float64 Floating numbers
		complex64, complex128 Complex numbers
*/
