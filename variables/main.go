package main

import "fmt"

func main() {
	var name string = "Flash" // type is string
	num1 := 10 // type is inferred
	var num2 = 20 // type is inferred
	fmt.Println(name)
	fmt.Println(num1)
	fmt.Println(num2)

	// Multiple variable declaration
	var one, two, three = 1,2,3
	fmt.Println(one)
	fmt.Println(two)
	fmt.Println(three)
}