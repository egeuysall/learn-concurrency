package main

import "fmt"

// square calculates the sum of squares of all elements in slice s
// and sends the result through channel c
func square(s []int, c chan int) {
	sum := 0

	// Calculate sum of squares
	for _, v := range s {
		sum += v * v
	}

	// Send the result to the channel
	c <- sum
}

func main() {
	// Create a slice of 20 integers
	s := make([]int, 20)
	// Fill slice with values 1 through 20
	for i := range 20 {
		s[i] = i + 1
	}

	// Create a channel that can pass integers
	c := make(chan int)

	// Start the square function in a goroutine (runs concurrently)
	go square(s, c)

	// Receive the result from the channel (this blocks until data is available)
	result := <-c

	fmt.Println(result)
}
