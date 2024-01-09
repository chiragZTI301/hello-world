package main

import "fmt"

func takeInput() {
	var firstname, lastName string
	var i, j int

	fmt.Scan(&firstname, &lastName)
	//read space separated values, and scan till mentioned values (or variables) couldn't get values.
	//like if we have press enter after entering 1 value then it will still active for getting remaining values.
	fmt.Println(firstname, lastName)

	fmt.Scanln(&firstname, &lastName) //does same but if we press enter then reemaining value will be assigned with 0 or "" or false
	fmt.Println(lastName, firstname)

	//yes we've another function called scanf,
	// receives the inputs and stores them based on the determined formats for its arguments.
	fmt.Print("Type two numbers: ")
	fmt.Scanf("%v%v", &i, &j)
	fmt.Println("Your numbers are:", i, "and", j)
}
