package main

import (
	"fmt"
	"time"
)

func main() {
	// Create two unbuffered channels for communication between goroutines
	c1 := make(chan int)
	c2 := make(chan int)

	// Launch first goroutine that sends values 1-5 to c1 (one per second)
	go func() {
		for i := 1; i <= 5; i++ {
			time.Sleep(1 * time.Second)
			c1 <- i
		}
		close(c1) // Always close channels when done sending
	}()

	// Launch second goroutine that sends values 10-50 to c2 (every 500ms)
	go func() {
		for i := 10; i <= 50; i += 1 {
			time.Sleep(500 * time.Millisecond)
			c2 <- i
		}
		close(c2) // Always close channels when done sending
	}()

	// Keep looping until both channels are nil (marked as closed and processed)
	for c1 != nil || c2 != nil {
		// SELECT statement: waits for the first channel that's ready to receive from
		// Key points about select:
		// - It blocks until one of its cases can proceed
		// - If multiple cases are ready, it picks one at random
		// - The default case runs immediately if no other case is ready (makes it non-blocking)
		select {
		// Case 1: Try to receive from c1
		case i, ok := <-c1:
			// ok is false when channel is closed and empty
			if !ok {
				c1 = nil // Set to nil so select ignores this case going forward
				continue
			}
			fmt.Println("Counter 1:", i)
		// Case 2: Try to receive from c2
		case i, ok := <-c2:
			// ok is false when channel is closed and empty
			if !ok {
				c2 = nil // Set to nil so select ignores this case going forward
				continue
			}
			fmt.Println("Counter 2:", i)
		// Default case: executes when no channels are ready
		// Without this, select would block waiting for a channel
		default:
			time.Sleep(50 * time.Millisecond)
		}
	}

	fmt.Println("Done!")
}
