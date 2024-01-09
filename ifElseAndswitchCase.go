package main

import (
	"fmt"
	"runtime"
	"time"
)

func switchCase() {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}
}

func switchCase1() {
	today := time.Now().Weekday()
	fmt.Println(today)
	switch time.Saturday {
	case today:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

func switchCase2(a, b []byte) int {
	for i := 0; i < len(a) && i < len(b); i++ {
		switch {
		case a[i] > b[i]:
			fmt.Println(i, a[i], b[i])
			return 1
		case a[i] < b[i]:
			fmt.Println(i, a[i], b[i])
			return -1
		// default:
		// 	return 3
		}
	}
	switch {
	case len(a) > len(b):
		fmt.Println(len(a), len(b))
		return 1
	case len(a) < len(b):
		fmt.Println(len(a), len(b))
		return -1
	}
	return 0
}
