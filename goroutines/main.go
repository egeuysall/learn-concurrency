package main

import (
	"fmt"
	"time"
)

// A function that prins the parameter it's given after a delay of 100ms.
func say(s string) {
	for range 5 {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world")
	say("hello")
}
