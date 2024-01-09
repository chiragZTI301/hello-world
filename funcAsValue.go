package main

import "fmt"

func functionAsValue() {
	hypot := func(x, y int) string {
		return "x + y"
	}

	fmt.Println(hypot(5, 12))
	fmt.Println(compute(hypot))
}

var abcd = 0

func compute(abc func(int, int) string) string {
	return abc(5, 12)
}
