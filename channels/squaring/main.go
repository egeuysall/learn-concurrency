package main

import "fmt"

func square(s []int, c chan int) {
	sum := 0

	for _, v := range s {
		sum += v * v
	}

	c <- sum
}

func main() {
	s := make([]int, 10)
	for i := range 10 {
		s[i] = i + 1
	}
	c := make(chan int)

	go square(s, c)
	result := <-c

	fmt.Println(result)
}
