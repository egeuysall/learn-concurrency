package main

import "fmt"

func square(x int, c chan int) {
	// range over x iterations (0 to x-1)
	for i := range x {
		// send the square of (i+1) to the channel
		// this won't block because we have a buffered channel with capacity 16
		c <- (i + 1) * (i + 1)
	}
	// close the channel to signal no more values will be sent
	// this is important - without close(), the range loop in main would block forever
	close(c)
}

func main() {
	// create a buffered channel with capacity 16
	// buffered channels allow sending values without blocking until the buffer is full
	ch := make(chan int, 16)

	// launch square function as a goroutine
	// it will run concurrently with the main function
	go square(16, ch)

	// range over the channel - this will receive values until the channel is closed
	// when the channel is closed and empty, the loop will exit
	for i := range ch {
		fmt.Println(i)
	}
}
