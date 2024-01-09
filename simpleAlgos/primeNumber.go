package simplealgos

import (
	"fmt"
)

func CheckPrimeNumber() {
	num := 2
	if(num == 2){
		fmt.Println("Prime it is.")
		return
	}
	for i := 2; i < num-1; i++ {
		if num%i == 0 {
			fmt.Println("Not Prime!")
			return
		}
	}
	fmt.Println("Prime it is.")
}
