package simplealgos

import (
	"fmt"
)

func BinarySearch() int {
	arr, x := [8]int{34, 35, 41, 42, 54, 56, 76, 90}, 91
	//we asumed that data is already sorted.

	l, r, m := 0, len(arr)-1, 0
	for l <= r {
		m = l + (r-l)/2

		if arr[m] == x {
			fmt.Println("Yes Matched.")
			return m
		}

		if arr[m] < x {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	fmt.Println("Not Matched.")
	return -1
}
