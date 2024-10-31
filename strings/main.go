package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hello"
	str1 := "Some1"
	str2 := "Some2"

	fmt.Printf("%c \n", str[0])
	fmt.Println(str[0:3])
	fmt.Println(strings.Compare(str1, str2))
	fmt.Println(strings.Contains(str, "a"))
	fmt.Println(strings.ToLower(str))
}