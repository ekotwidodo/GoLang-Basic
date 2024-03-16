package main

import "fmt"

const LENGTH int = 10
const WIDTH = 5

func main() {
	const firstName string = "Eko"
	const lastName = "Teguh"

	fmt.Println(firstName)
	fmt.Println(lastName)

	// error
	// firstName = "Tidak bisa diubah"
	// lastName = "Tidak bisa diubah"

	fmt.Println(LENGTH)
	fmt.Println(WIDTH)
}
