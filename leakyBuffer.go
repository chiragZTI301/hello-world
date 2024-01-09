package main

import (
	"fmt"
	"time"
)

func leakyBuffer1() {

}

func leakyBuffer2() {
	bufferedChannel := make(chan int, 2)

	// Stop signal channel
	stop := make(chan struct{})

	// Sender: A goroutine continuously sending values to the channel
	go func() {
		defer close(bufferedChannel)
		for i := 0; ; i++ {
			select {
			case bufferedChannel <- i:
			case <-stop:
				return
			}
			time.Sleep(500 * time.Millisecond) // Shorter sleep duration
		}
	}()

	// Receiver: A goroutine that occasionally receives values from the channel
	go func() {
		defer close(stop)
		for {
			select {
			case value, ok := <-bufferedChannel:
				if !ok {
					fmt.Println("Channel closed")
					return
				}
				fmt.Println("Received:", value)
				// Simulate some processing
				time.Sleep(2 * time.Second)
			default:
				// No value ready to be received
				fmt.Println("No value ready")
				time.Sleep(time.Second)
			}
		}
	}()

	// Let the program run for a while
	time.Sleep(10 * time.Second)
}
