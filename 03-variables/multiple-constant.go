package main

import "fmt"

const (
	X int = 10
	Y     = 3.14
	Z     = "Hello World!"
)

func main() {
	const (
		firstName string = "Eko"
		lastName         = "Teguh"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)

	// error
	// firstName = "Tidak bisa diubah"
	// lastName = "Tidak bisa diubah"

	fmt.Println(X)
	fmt.Println(Y)
	fmt.Println(Z)
}
