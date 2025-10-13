package main

import (
	"fmt"
	"sync"
	"time"
)

// worker is a goroutine that processes tasks from a channel
// id: unique identifier for this worker
// tasks: receive-only channel from which tasks are consumed
// wg: WaitGroup to signal when this worker is done
func worker(id int, tasks <-chan string, wg *sync.WaitGroup) {
	// Signal the WaitGroup that this worker is done when the function exits
	defer wg.Done()

	for {
		select {
		// Try to receive a task from the tasks channel
		case task, ok := <-tasks:
			// If the channel is closed, exit the worker
			if !ok {
				return
			}
			// Process the task (simulated with a sleep)
			fmt.Println("Worker", id, "processing task", task)
			time.Sleep(200 * time.Millisecond)

		// If no task arrives within 1 second, print idle message
		case <-time.After(time.Second):
			fmt.Println("Worker", id, "is idle")

		// If no task is immediately available, sleep briefly to avoid busy-waiting
		default:
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func main() {
	// Create a channel for sending tasks to workers
	tasks := make(chan string)

	// WaitGroup to wait for all workers to finish
	var wg sync.WaitGroup

	// Create and start 5 worker goroutines
	numWorkers := 5
	for i := range numWorkers {
		wg.Add(1) // Increment the WaitGroup counter for each worker
		go worker(i, tasks, &wg)
	}

	// Create a ticker that fires every 500 milliseconds
	ticker := time.NewTicker(500 * time.Millisecond)

	// Start a goroutine to send 10 tasks at regular intervals
	go func() {
		for i := range 10 {
			<-ticker.C                         // Wait for the next tick
			tasks <- fmt.Sprintf("Task-%d", i) // Send a task to the channel
		}
		// Close the tasks channel after all tasks are sent
		// This signals workers to exit
		close(tasks)
	}()

	// Wait for all workers to complete
	wg.Wait()
	fmt.Println("All tasks are done")
}
