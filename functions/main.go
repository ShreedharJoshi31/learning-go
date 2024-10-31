package main

import "fmt"

func printName(name string) {
	fmt.Println(name)
}

func showNumbers(numbers ...int) { // By using ... we can take the parameters infinte times it is called as an unlimited parameter
	fmt.Println(numbers)
}

func showUsers(admin string ,names ...string) { // The unlimited parameter should always be at the end
	fmt.Println("The admin is", admin)
	for idx, val := range names {
		fmt.Println(idx, ":", val)
	}
}

func add(x int, y int) int {
	return x + y
}

func main() {
	printName("Flash")
	showNumbers(1,2,3,4,5,6,7,8,9)
	showUsers("Harry", "Hagrid", "Dumbledore")
	fmt.Println(add(1, 2))

	// Function expression
	addFunc := func(x int, y int) int {
		return x + y
	}
	sum := addFunc(1, 2)
	fmt.Println(sum)

	// Anonymous funtion
	func () {
		fmt.Println("HELLO")
	}()
	subtract := func (x int, y int) int {
		return x - y
	} (10, 5)
	fmt.Println(subtract)
}