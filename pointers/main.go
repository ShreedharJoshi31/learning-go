package main

import "fmt"

func main() {
	m := 10
	fmt.Println(&m)
	var pointerToM *int = &m
	fmt.Println(pointerToM)
	fmt.Println(*pointerToM) // Value of m
	*pointerToM = 20 // This will also change the value of m
	fmt.Println(m)
}