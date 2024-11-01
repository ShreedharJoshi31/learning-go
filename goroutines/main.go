package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := range 10 {
		time.Sleep(200 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

func main() {
	go printNumbers()
}