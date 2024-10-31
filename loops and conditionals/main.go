package main

import "fmt"

func main() {
	for num := 1; num <= 30; num++ {
		if num == 3 || num == 15 {
			continue
		}
		if num == 25 {
			break
		}
		if num % 3 == 0 {
			fmt.Println("FIZZ")
		} else if num % 5 == 0 {
			fmt.Println("BUZZ")
		} else {
			fmt.Println(num)
		}
	}

	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}
}