package main

import "fmt"

// multiply calculates the product of all integers in slice s
// and sends the result through channel c
func multiply(s []int, c chan int) {
	sum := 1

	// Calculate product of all values in the slice
	for _, v := range s {
		sum *= v
	}

	// Send the result to the channel
	// This is a "send" operation: c <- value
	c <- sum
}

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Create a channel that can carry int values
	// Channels are used for communication between goroutines
	c := make(chan int)

	// Start the multiply function as a goroutine (concurrent execution)
	// The "go" keyword runs the function in a separate goroutine
	go multiply(s, c)

	// Receive from the channel - this blocks until a value is sent
	// This is a "receive" operation: <-c
	result := <-c

	fmt.Println(result)
}
