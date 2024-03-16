package main

import "fmt"

var x int = 1

// y := 2 // Error, tidak bisa digunakan di luar fungsi

func main() {
	var name string

	name = "Eko Teguh"
	fmt.Println(name)

	name = "Eko Widodo"
	fmt.Println(name)

	height := 175
	fmt.Println(height)

	var a string
	var b int
	var c bool

	var d, e, f, g = 1, 2, 3, "Hello World"

	fmt.Println(x)
	// fmt.Println(y)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)

}
