package main

import (
	"fmt"
	_ "io"
	//here _ "io" means it won't be used explicitly, but imported for side effects or side usage.
	// For example, during its init function, the net/http/pprof package registers HTTP handlers that provide debugging information.
	// It has an exported API, but most clients need only the handler registration and access the data through a web page.
)

// func init() {
// 	fmt.Println("Init Function Just Called.")
// 	callFunction(1)
// }

func callFunction(i int) {
	var a int
	var b float64
	a = 2
	b = 3.2
	fmt.Println("Init Function has been called already.", i)
	fmt.Println(float64(a) * b)
	//in int * float64 you'll need to convert int into float64.
}
