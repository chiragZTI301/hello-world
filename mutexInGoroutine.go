package main

// A Mutex is a method used as a locking mechanism to ensure that only one Goroutine is accessing the critical section of code at any point of time
// This is done to prevent race conditions from happening. Sync package contains the Mutex. Two methods defined on Mutex => Lock and Unlock

import (
	"fmt"
	"sync"
)

var counter = 0
var mu sync.Mutex

func increment(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 10; i++ {
		mu.Lock()
		counter++
		mu.Unlock()
	}
}

func callMutex() {
	var wg sync.WaitGroup

	wg.Add(2)
	go increment(&wg)
	go increment(&wg)

	wg.Wait()

	fmt.Println("Counter:", counter)

	//s simply it is used to prevent race condition, for better example or code go to geeksforgeeks
	//what race condition meant: so whenever two or more threads access particular data (read or write) then data might be corrputed like data could be wrong placed there.
}
