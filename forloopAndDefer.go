package main

import (
	"fmt"
)

func loop() {
	for i := 0; i < 5; i++ {
		fmt.Println(i + 1)
	}
}

func loop1() {
	sum := 1
	for sum < 4 {
		sum += 1
		fmt.Println(sum)
	}
}

var a = [3]int{1,2,3}

func forlooprange() {
	for _, v := range a {
		fmt.Println(v)
	}
}

func calcSqrt(x float64) float64 {
	var t float64
	var sqrtroot = x / 2
	for t-sqrtroot != 0 {
		t = sqrtroot
		sqrtroot = (t + (x / t)) / 2
		fmt.Println("yes", sqrtroot, t)
	}
	return sqrtroot
}

func deferInGO() {
	fmt.Println("world")

	defer fmt.Println("world1")

	fmt.Println("hello")

	defer fmt.Println("world2")

	fmt.Println("hello1")
}
