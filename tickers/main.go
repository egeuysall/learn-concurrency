package main

import (
	"fmt"
	"time"
)

func main() {
	// Create a new ticker that will send a value on its channel every 1 second
	// A ticker is useful for executing code at regular intervals
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop() // always stop tickers when done to release resources

	// Loop 25 times
	for range 25 {
		// Block until the ticker sends a value on its C channel
		// This happens once per second based on the ticker's interval
		<-ticker.C
		fmt.Println("Tick at", time.Now())
	}
}
