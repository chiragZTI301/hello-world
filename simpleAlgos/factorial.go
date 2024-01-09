package simplealgos

import (
	"fmt"
)

func CalcFactorial() {
	num, factorial := 0, 1

	for k := num; k > 1; k-- {
		factorial *= k
	}

	fmt.Println(factorial)
}

func CalcFactorialRecusrion(num int) int {
	if num == 0 || num == 1 {
		return 1
	}

	return num * CalcFactorialRecusrion(num-1)
}
