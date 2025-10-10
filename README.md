# learn-concurrency

Hands-on Go concurrency examples using goroutines and channels.

## Goal

Learn concurrency patterns and parallel execution in Go.

## Mini Project Idea: Concurrent Task Processor

### Overview

Build a small Go program that simulates a SaaS background job processor. The program handles multiple tasks concurrently and demonstrates all the concurrency concepts you've learned.

### Features

- Multiple worker goroutines processing tasks concurrently
- Task queue implemented with a channel
- Workers use `select` to either:
  - Receive tasks from the channel
  - Time out if a task takes too long
  - Skip iteration if no task is ready (`default`)
- Periodic tasks added via a `time.Ticker` (e.g., cleanup jobs)
- WaitGroup ensures the main program waits for all workers to finish
- Task processing can have simulated timeouts

### Example Flow

1. Start 3 workers.
2. Add tasks to the queue every 2 seconds using a ticker.
3. Workers process tasks, respecting a timeout.
4. WaitGroup ensures main waits until all tasks are processed.
5. Program ends when all tasks are done.
