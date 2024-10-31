package main

import "fmt"

func main() {
	userInfo := map[string]int{
		"Flash": 21,
		"Zoom":  22,
		"Harry": 18,
	}

	fmt.Println(userInfo)
	userInfo["Ron"] = 20

	fmt.Println(userInfo["Flash"])

	for key, value := range userInfo {
		fmt.Printf("%s : %d\n", key, value)
	}
}