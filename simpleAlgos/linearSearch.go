package simplealgos

import (
	"fmt"
)

func LinearSearch() {
	arr, num := [8]int{54, 76, 41, 35, 42, 34, 56, 90}, 43

	for i := 0; i < len(arr); i++ {
		if arr[i] == num {
			fmt.Println("yes, number matched.")
			return
		}
	}
	fmt.Println("Number didn't match.")
}
