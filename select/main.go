package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		// This case sends the current fibonacci number (x) to channel c
		// If the receiver is ready to read from c, this case executes
		// Then we calculate the next fibonacci numbers
		case c <- x:
			x, y = y, x+y
		// This case listens for a signal on the quit channel
		// When something is sent to quit, this case executes
		// and the function returns, ending the fibonacci generation
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for range 10 {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
