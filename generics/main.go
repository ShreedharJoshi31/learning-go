package main

import "fmt"

func printItem[T any](item1 T, item2 T) (T, T) {
	return item1, item2
}

func anotherPrintItem[T1 string | int, T2 bool | int](item1 T1, item2 T2) (T1, T2) {
	return item1, item2
}

func printArray[T any](arr []T) {
	for _, value := range arr {
		fmt.Println(value)
	}
}

func main() {
	num1, num2 := printItem(10, 20)
	str1, str2 := printItem("One", "Two")
	bool1, bool2 := printItem[bool](true, false)

	item1, item2 := anotherPrintItem[string, bool]("Hello", true)

	fmt.Println(num1, num2)
	fmt.Println(str1, str2)
	fmt.Println(bool1, bool2)
	fmt.Println(item1, item2)

	numberArray := []int{1, 2, 3, 4, 5}
	stringArray := []string{"One", "Two", "Three"}

	printArray(numberArray)
	printArray(stringArray)
}