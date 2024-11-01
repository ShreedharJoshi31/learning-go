// select is basically a switch case for channels in go

package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 :=  make(chan string)
	ch2 :=  make(chan string)

	go func() {
		time.Sleep(4 * time.Second)
		ch1 <- "Message from channel 1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Message from channel 2"
	}()

	select{
	case msg := <-ch1:
		fmt.Println("Received from channel 1:", msg)
	case msg := <-ch2:
		fmt.Println("Received from channel 2:", msg)
	case <-time.After(5*time.Second):
		fmt.Println("No comms")
	}
}