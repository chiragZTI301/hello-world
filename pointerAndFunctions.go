package main

import "fmt"

func normal() {
	var a = person{"Chirag", "20"}
	updateFunc1(&a)
	fmt.Println(a)
}

func updateFunc1(a *person) {
	a.age = "21"
	fmt.Println(a)
}

// same thing, can not use *int or *string or *float64 etc


//Note
// There are two reasons to use a pointer receiver.

// The first is so that the method can modify the value that its receiver points to.

// The second is to avoid copying the value on each method call. This can be more efficient if the receiver is a large struct, for example.