package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // decrements counter when function finishes
	time.Sleep(time.Second)
	fmt.Println("Worker", id, "done")
}

func main() {
	var wg sync.WaitGroup
	numWorkers := 3

	wg.Add(numWorkers) // number of goroutines to wait for

	for i := 1; i <= numWorkers; i++ {
		go worker(i, &wg)
	}

	wg.Wait() // block until all workers call Done()
	fmt.Println("All workers finished")
}
