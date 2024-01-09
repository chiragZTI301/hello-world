package main

import "fmt"

func arrays() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	b := [2]int{2, 3}
	fmt.Println(b)
}

func multiDimensionalArray() {
	var s = [2][2]string{{"Chirag", "Makwana"},{"Dhyey","Panchal"}}
	fmt.Println(s)

	var a [2][2]int
	fmt.Println(a)
	a[0][1] = 1;
	fmt.Println(a)

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Println(s[i][j])
		}
	}
}

func slices() {
	a := []int{1, 2, 3}
	fmt.Println(a)

	array := [5]int{1, 2, 3, 4, 5}
	a = array[1:4]
	fmt.Println(a)
}

func sliceReferToArray() {
	a := [5]int{1, 2, 3, 4, 5}
	b := a[1:2]
	b[0] = 10
	fmt.Println(a)
	fmt.Println(b)
}

func sliceDefaults() {
	// The default is zero for the low bound and the length of the slice for the high bound.

	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Printf("%d", cap(s))

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}

func sliceUsingMakeFunc() {
	s := make([]int, 4, 7) //if you write make([]int, 5) then it wil be considered as length and capacity both are 5
	fmt.Println(s[0], s)
}

func iterateSlice() {
	s := []int{1,2,3,4,5,6,7}	
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}

	for index,ele := range s { //and if you don't want index or range then you can put _ there. and also you can use like for index := range but not for ele.
		fmt.Println(index, ele)
	}
}

func sliceOfSlices() {
	a := [][]string{
		{"abc"},
		{"def"},
		{"ghi"},
		{"jkl", "mno", "pqr"},
	}
	fmt.Println(a)
}

func appendToSlice() {
	s := []int{1,2,3}
	s = append(s, 4)
	fmt.Println(s)
	s = append(s, 5)
	fmt.Println(s)
}