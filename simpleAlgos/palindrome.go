package simplealgos

import (
	"fmt"
	"strings"
)

func CheckPalindromeString() {
	s1 := "nayAn"
	s1 = strings.ToLower(s1)

	for i, j := 0, len(s1)-1; i < len(s1); i, j = i+1, j-1 {
		if s1[i] != s1[j] {
			fmt.Println("Not Palindrome!")
			return
		}
	}
	fmt.Println("Palindrome it is.")
}

func CheckPalindromeNumber() {
	num := 121
	input_num := num
	var remainder int
	res := 0
	for num>0 {
	   remainder = num % 10
	   res = (res * 10) + remainder
	   num = num / 10
	}
	if input_num == res {
		fmt.Println("Palindrome it is.")
	} else {
		fmt.Println("Not Palindrome!")
	}
}
