package main

import (
	"fmt"
)

func isContains[T comparable](s []T, x T) int {
	//here [T any] could be there but then we can not compare v with x at line no of 10
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

func compare() {
	si := []int{10, 20, 15, -10}
	fmt.Println(isContains(si, 15))

	ss := []string{"foo", "bar", "baz"}
	fmt.Println(isContains(ss, "foo"))
}
