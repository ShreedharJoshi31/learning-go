package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}
	fmt.Println(len(numbers))
	fmt.Println(cap(numbers))
	numbers = append(numbers, 3, 4, 5)
	fmt.Println(numbers)

	// make() function to create slices takes 3 positional args type, lenght, capacity
	slice := make([]int, 3, 5)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	fmt.Println(slice)
	slice = append(slice, 1, 2, 3, 4, 5, 6)
	slice[0] = 15
	fmt.Println(slice)
}