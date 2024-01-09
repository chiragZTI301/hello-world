package main

import (
	"fmt"
	"time"
)

func printForGorutine() {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(i + 1)
	}
}

func callGorutine() {
	go printForGorutine()

	//here by default main goroutine is working and if you hit go then another goroutine will start executing.
	//remeber that all gorutines will stop executing after main goroutine stop executing (means completing it execution)
	//to ensure all goruitnes complete their execution before main one, we need to use chanels and wait.syncGroup etc.

	for i := 0; i < 5; i++ {
		time.Sleep(200 * time.Millisecond)
		fmt.Println("Main:", i+1)
	}
}
