package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)

	go func() {
		time.Sleep(2 * time.Second) // simulates long task
		c <- 42
	}()

	// select statement waits for whichever case becomes ready first
	select {
	case val := <-c:
		// This case triggers if the goroutine sends a value before timeout
		fmt.Println("Received:", val)
	case <-time.After(1 * time.Second):
		// time.After() returns a channel that sends a value after the specified duration
		// This case triggers if 1 second passes before receiving from c
		// Since the goroutine sleeps for 2 seconds, this timeout will always fire first
		fmt.Println("Timeout! Task took too long.")
	}
}
