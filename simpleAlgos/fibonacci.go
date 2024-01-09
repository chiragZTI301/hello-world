package simplealgos

import "fmt"

func CalculateFibonacci() {
	num := 8

	i, j, v := 0, 1, 1

	for k := 2; k < num; k++ {
		v = i + j
		i = j
		j = v
		fmt.Println(i,j)
	}

	fmt.Println(v)
}