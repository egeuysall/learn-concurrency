package main

import "fmt"

func multiply(s []int, c chan int) {
	sum := 1

	for _, v := range s {
		sum *= v
	}

	c <- sum
}

func main() {
	s := make([]int, 10)
	for i := range 10 {
		s[i] = i + 1
	}
	c := make(chan int)

	go multiply(s, c)
	result := <-c

	fmt.Println(result)
}
