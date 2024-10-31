package main

import "fmt"

func main() {
	var arr = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)

	var arr1 [5]int
	arr1[1] = 10
	fmt.Println(arr1)

	fmt.Println(len(arr1))
	fmt.Println(cap(arr1))

	for idx, number := range arr {
		fmt.Println(idx, number)
	}
}