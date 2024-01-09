package main

import "fmt"

type person struct {
	name, age string
}

//yes you can declare method for non strucy type value just like below,
type number int

func func1() {
	var a = person{"Chirag", "20"}
	updateFunc(a)
	fmt.Println(a)
	a.updateMethod()
	fmt.Println(a)
	//another point to be noticed is that we can assign var a = &person{"Chirag", "20"} as well.

	var b = number(6044)
	b.updateMethod1()
}

func updateFunc(a person) {
	a.age = "21"
}

func (p *person) updateMethod() {
	p.age = "21"
}

func (n number) updateMethod1() {
	fmt.Println(n)
	// so here you can not take *number as per official doc of go.
	//This means the receiver type has the literal syntax *T for some type T. (Also, T cannot itself be a pointer such as *int.)
	//so what is slution => try to return it.
}
