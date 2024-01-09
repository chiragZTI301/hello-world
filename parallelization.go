package main

import (
	"fmt"
	"runtime"
	"sync"
)

func printNumbers(start, end int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := start; i <= end; i++ {
		fmt.Println(i)
	}
}

func parallelizeFunction()  {
	numCPU := runtime.NumCPU()

	// Set GOMAXPROCS to the number of CPU cores
	runtime.GOMAXPROCS(numCPU)

	var wg sync.WaitGroup

	// Divide the work among multiple goroutines
	for i := 0; i < numCPU; i++ {
		wg.Add(1)
		go printNumbers(i*10+1, (i+1)*10, &wg)
	}

	// Wait for all goroutines to finish
	wg.Wait()
}

func getCPUCores() {
	cpu := runtime.NumCPU()
	fmt.Println(cpu)
}