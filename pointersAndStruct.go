package main

import (
	"fmt"
)

func pointers() {
	i, j := 1, 2
	p, q := &i, &j
	fmt.Println(p, q, *p, *q)
	//you can set value of i and j by using *p and *q as well now...
}

type address struct {
	name, city, state string
	//and if you don't want to use semi colon then simply write it in next line, just like below
	pincode int
}

func structures() {
	var a address
	// a.name = "Chirag"
	// a.city = "Ahmedabad"
	// a.state = "Gujarat"
	// a.pincode = 380061

	a = address{city: "Ahmedabad", name: "Chirag", state: "Gujarat", pincode: 380061}
	fmt.Println(a) //you can access it by using dot (.) as well
}

func structAndPointers() {
	a := address{city: "Ahmedabad", name: "Chirag", state: "Gujarat", pincode: 380061}
	b := &a
	fmt.Println(b.city) //you can do it by using *b.city as well
}
