package main

import "fmt"

func sendData(ch chan<- int) {
	ch <- 20
}

func main() {
	ch := make(chan int)
	go sendData(ch)
	value := <-ch
	fmt.Println(value)

	// Unbuffered channel
	unbufferedCh := make(chan int)

	// Buffered channel
	bufferedCh := make(chan string, 2)

	go func() {
		unbufferedCh <- 20
	}()

	go func() {
		bufferedCh <- "Hello"
		bufferedCh <- "World"
		close(bufferedCh)
	}()

	unbufferedValue := <-unbufferedCh
	fmt.Println(unbufferedValue)

	bufferedValue1 := <-bufferedCh
	bufferedValue2 := <-bufferedCh

	fmt.Println(bufferedValue1)
	fmt.Println(bufferedValue2)

	bufferedCh2 := make(chan int, 5)
	go func() {
		bufferedCh2 <- 10
		bufferedCh2 <- 20
		bufferedCh2 <- 30
		bufferedCh2 <- 40
		close(bufferedCh2)
	}()

	for val := range bufferedCh2 {
		fmt.Println(val)
	}


}