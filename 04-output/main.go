package main

import "fmt"

func main() {
	var firstName, middleName, lastName string = "Eko", "Teguh", "Widodo"
	var a, b = 100, 200

	var nickName, myAddress = "Eko", "Bandar Lampung"

	var x string = "Halo"
	var y int = 25
	var angka = 20.5
	var txt = "Hello World!"

	fmt.Print(firstName, " ", middleName, " ", lastName, "\n")
	fmt.Print(a, b, "\n")

	fmt.Println(nickName, myAddress)

	fmt.Printf("nilai x adalah : %v dan tipe data x adalah : %T\n", x, x)
	fmt.Printf("nilai y adalah : %v dan tipe data x adalah : %T\n", y, y)
	fmt.Printf("%v%%\n", angka)
	fmt.Printf("%#v", txt)
}
