package main

import (
	"fmt"
	"time"
)

// The select statement in Go is used to handle multiple communication operations simultaneously.
// It provides a way to wait on multiple communication operations, such as sending or receiving data from channels, and selects the first operation that's ready to proceed.
// If multiple operations are ready, one is chosen at random.

func selectStatement() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch1 <- "Message from ch1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "Message from ch2"
	}()

	// select {
	// case msg1 := <-ch1:
	// 	fmt.Println("Received from ch1:", msg1)
	// case msg2 := <-ch2:
	// 	fmt.Println("Received from ch2:", msg2)
	// case <-time.After(3 * time.Second):
	// 	fmt.Println("Timeout: No communication within 3 seconds")
	// }

	for {
		select {
		case msg1 := <-ch1:
			fmt.Println("Received from ch1:", msg1)
		case msg2 := <-ch2:
			fmt.Println("Received from ch2:", msg2)
		case <-time.After(5 * time.Second):
			fmt.Println("Timeout: No communication within 3 seconds")
			return
		}
	}
}
