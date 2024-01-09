package main

import (
	"fmt"
	"math"
)

// import "fmt"

type shape interface {
	area() int
}

type square struct {
	height int
}

type circle struct {
	radius int
}

func (s square) area() int {
	return s.height * s.height
}

func (c circle) area() int {
	return int(float64(c.radius) * float64(c.radius) * math.Pi)
}

func calcArea(sh shape) {
	fmt.Println(sh.area())
}

func InterfaceCheck(s shape) {
	if sq, ok := s.(square); ok {
		fmt.Printf("Square Area: %d\n", sq.area())
	} else {
		fmt.Println("Not a square.")
	}

	if cir, ok := s.(circle); ok {
		fmt.Printf("Circle Area: %d\n", cir.area())
	} else {
		fmt.Println("Not a circle.")
	}
}

func callInterfaceCheck() {
	c := circle{4}
	// s := square{4}
	InterfaceCheck(c)
}

func calcArea1() {
	c := circle{4}
	// c1 := &circle{4}
	s := square{4}
	calcArea(c)
	calcArea(s)
	fmt.Printf("%v %T", c, c)
}

func emptyInterface() {
	var c interface{}
	fmt.Printf("%v %T", c, c)

	c = 42
	fmt.Printf("%v %T", c, c)

	c = "Chirag Makwana"
	fmt.Printf("%v %T", c, c)

	var s = square{4}
	c = s
	fmt.Printf("%v %T", c, c)
}

func typeassertion() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // panic
	fmt.Println(f)
}

func typeSwitches(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("%v %T", v, v)
	case string:
		fmt.Printf("%v %T", v, v)
	case float64:
		fmt.Printf("%v %T", v, v)
	default:
		fmt.Println("It is in default seciton")
	}
}

type Person struct {
    FirstName string
    LastName  string
}

func (p Person) String() string {
    return fmt.Sprintf("%s %s", p.FirstName, p.LastName)
}

func stringerFunc()  {
	// The fmt package look for this interface to print values.
	var p = Person{"Sameer", "Makwana"}
	fmt.Println(p)
}