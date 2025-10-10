package main

import "fmt"

func main() {
	// Create a buffered channel with capacity of 2
	// Buffered channels can hold values without blocking until the buffer is full
	ch := make(chan int, 2)

	// Send values to the channel - these won't block because buffer has space
	// First value goes into the buffer
	ch <- 1
	// Second value goes into the buffer (now buffer is full)
	ch <- 2

	// Receive values from the channel - FIFO (First In, First Out) order
	// This receives the first value sent (1) and removes it from buffer
	fmt.Println(<-ch)
	// This receives the second value sent (2) and removes it from buffer
	fmt.Println(<-ch)
}
