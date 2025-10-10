package main

import (
	"fmt"
	"time"
)

// A function that prins the parameter it's given after a delay of 100ms.
func say(s string) {
	// Loop 5 times
	for range 5 {
		// Wait 100 milliseconds before printing
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	// Start a goroutine - this runs concurrently in the background
	// The "go" keyword creates a new goroutine that executes say("world")
	go say("world")

	// This runs in the main goroutine (synchronously)
	// The program will wait for this to complete before exiting
	say("hello")

	// Note: The program may exit before the goroutine finishes
	// since main() doesn't wait for the goroutine to complete
}
