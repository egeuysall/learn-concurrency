package main

import (
	"fmt"
	"sync"
	"time"
)

func sendNotification(userID int, wg *sync.WaitGroup) {
	defer wg.Done()

	// Simulate variable time to send notification
	time.Sleep(time.Duration(userID%3+1) * 500 * time.Millisecond)
	fmt.Println("Notification sent to user", userID)
}

func main() {
	users := []int{101, 102, 103, 104, 105} // could be any number
	var wg sync.WaitGroup

	for _, user := range users {
		wg.Add(1)
		go sendNotification(user, &wg)
	}

	done := make(chan struct{})
	go func() {
		wg.Wait()       // wait for all notifications to finish
		close(done)     // signal completion
	}()

	select {
	case <-done:
		fmt.Println("All notifications sent successfully!")
	case <-time.After(2 * time.Second):
		fmt.Println("Timeout! Some notifications took too long.")
	}
}
