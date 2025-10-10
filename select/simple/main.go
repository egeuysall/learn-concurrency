package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Hello from ch1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Hello from ch2"
	}()

	// Loop twice to receive messages from both channels
	for range 2 {
		// select statement allows us to wait on multiple channel operations
		// It will block until one of the channels has data ready to receive
		// Whichever channel receives data first will be processed
		select {
		case msg1 := <-ch1:
			// This case executes when ch1 receives a message (after ~1 second)
			fmt.Println(msg1)
		case msg2 := <-ch2:
			// This case executes when ch2 receives a message (after ~2 seconds)
			fmt.Println(msg2)
		}
	}
}
