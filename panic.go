package main

import (
	"fmt"
)

func panicFunction() {
	fmt.Println("Start of the program")

    // Simulate a panic situation
    simulatePanic()

    fmt.Println("End of the program")
}

func simulatePanic()  {
	defer func() {
		if r := recover(); r != nil {
            fmt.Println("Recovered from panic:", r)
        }
    }()

    fmt.Println("About to panic")
    panic("Panic situation")
    // unreachable code: fmt.Println("This line won't be executed")
}