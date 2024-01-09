package main

import (
	"fmt"
	"sync"
)

func createChannel() {
	// Create a channel for string messages
	messageChannel := make(chan string) // we can have buffered channel as well like make(chan string, 100)

	// Use a WaitGroup to wait for goroutines to finish
	var wg sync.WaitGroup

	// Launch two sender goroutines
	wg.Add(2)
	go sender("Hello", messageChannel, &wg)
	go sender("World", messageChannel, &wg)

	go func() {
		// Wait for both sender goroutines to finish
		wg.Wait()

		// Close the channel since no more messages will be sent
		close(messageChannel)
	}()

	// Receive and print messages from the channel
	for message := range messageChannel {
		fmt.Println("Received:", message)
	}
}

func sender(message string, ch chan<- string, wg *sync.WaitGroup) {
	// defer wg.Done(): This is deferred until the surrounding function completes. wg.Done() decrements the WaitGroup counter by 1,
	// indicating that one of the goroutines has finished its execution.
	defer wg.Done()

	// Send the message to the channel
	ch <- message
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func callFibonacci() {
	c := make(chan int, 10)
	fmt.Println(cap(c))
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}
