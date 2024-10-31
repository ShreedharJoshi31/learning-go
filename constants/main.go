package main

import "fmt"

const number = 20 // Cannot be changed

func main() {
	const word string = "Hello" // Cannot be changed
	fmt.Println(number)
	fmt.Println(word)
}