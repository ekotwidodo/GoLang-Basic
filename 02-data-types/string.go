package main

import "fmt"

func main() {
	fmt.Println("Eko")
	fmt.Println("Eko Teguh")
	fmt.Println("Eko Teguh Widodo")

	fmt.Println(len("Eko"))
	fmt.Println("Eko Teguh"[0])        // result = 69 (byte karakter E)
	fmt.Println("Eko Teguh Widodo"[1]) // result = 107 (byte karakter k)
}
