package main

import (
	"fmt"
)

type Person1 struct {
	FirstName string
	LastName  string
}

type Person2 struct {
	Age int
}

type Employee struct {
	person Person1
	persn2 *Person2
	empId  int
}

func changePersonAge(p *Person2) {
	p.Age = 21
}

func embeddingFunc() {
	employee := Employee{
		person: Person1{FirstName: "Chirag", LastName: "Makwana"},
		persn2: &Person2{Age: 20},
		empId:  1,
	}

	p2 := employee.persn2

	fmt.Println(employee.persn2.Age)
	changePersonAge(p2)
	fmt.Println(employee.persn2.Age)
}