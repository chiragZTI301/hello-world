package simplealgos

import (
	"fmt"
)

func BubbleSort() {
	arr := []int{54, 76, 41, 35, 42, 34, 56, 90}
	var i, j, k int

	for i = 0; i < len(arr)-1; i++ {
		for j = 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				k = arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = k
			}
		}
	}

	fmt.Println(arr)
}
