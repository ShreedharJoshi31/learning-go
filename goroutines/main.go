package main

import (
	"fmt"
	"sync"
	"time"
)

func printNumbers(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range 10 {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2) // Add how many goroutines we have to perform
	go printNumbers(&wg)
	go printNumbers(&wg)
	wg.Wait() // This will wait until all the goroutines are executed
}