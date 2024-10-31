package main

import "fmt"

// If a function is passed to another function as an argument.

func addName(name string, age int, printName func(string, int)) {
	printName(name, age)
}

func main() {
	addName("Flash", 10, func(nm string, a int) {
		fmt.Println(nm, a)
	})
}