package main

import (
	"fmt"

	// dir "golearning/dir"
	// dir2 "golearning/dir/dir2"
	// "golearning/dir/dir3"
	// "golearning/dir1"
	//this illustarte how directory and it's sub-directory being used whenever they uses same or different packages.

	simplealgos "golearning/simpleAlgos"
	"os"
	// "strconv"
)

var abc int = 1

func main() {
	// dir.Func1()
	// dir2.Func3()
	// dir4.Func4()
	// dir5.Func5()

	// simplealgos.CheckPalindromeString()
	// simplealgos.CheckPalindromeNumber()
	// simplealgos.CheckPrimeNumber()
	// simplealgos.CalculateFibonacci()
	// simplealgos.CalcFactorial()
	// fmt.Println(simplealgos.CalcFactorialRecusrion(2))
	// simplealgos.LinearSearch()
	// simplealgos.BubbleSort()
	// simplealgos.BinarySearch()
	fmt.Println(simplealgos.GCD(648, 172))

	// a := 2
	// fmt.Println("Hello World")
	// fmt.Printf("Type of a is %T", a)
	// fmt.Printf("Type of strconv.Itoa(a) is %T", strconv.Itoa(a))
	// fmt.Println("hello world ", a)
	// fmt.Println(main2())
	// dir.Func2()
	// main1()
	// helloMain(1, 2, "3")

	// forlooprange()
	// ifcond()
	// fmt.Println(calcSqrt(25))
	// switchCase()
	// switchCase1()
	// result := switchCase2([]byte("aaa"), []byte("aaa"))
	// fmt.Println(result)
	// deferInGO()

	// pointers()
	// structures()
	// structAndPointers()
	// arrays()
	// multiDimensionalArray()
	// slices()
	// sliceReferToArray()
	// sliceDefaults()
	// sliceUsingMakeFunc()
	// iterateSlice()
	// sliceOfSlices()
	// appendToSlice()
	// mapBasics()
	// functionAsValue()
	// func1()
	// normal()
	// calcArea1()
	// emptyInterface()
	// typeassertion()
	// typeSwitches(21.4)
	// stringerFunc()
	// err()
	// readData()
	// readDataFromFile()
	// takeInput()
	// img()
	// compare()
	// callGorutine()
	// createChannel()
	// callFibonacci()
	// selectStatement()
	// callMutex()
	// callFunction(0)
	// callInterfaceCheck()
	// embeddingFunc()
	// getCPUCores()
	// parallelizeFunction()
	// panicFunction()
	// leakyBuffer2()

	//i am here writing this comment to test git rebase command.
	//i am here writing this comment to test git rebase command.
}

func helloMain(x, y int, z string) {
	fmt.Printf(" type of x, y, z is %T, %T, %T", x, y, z)
}

func main1() {
	homeDir := os.Getenv("HOMEPATH")
	gopath := os.Getenv("GOPATH")
	fmt.Println("Home Directory:", homeDir)
	fmt.Println("GO Path:", gopath)

	/* envVars := os.Environ()
	for _, envVar := range envVars {
		fmt.Println(envVar)
	} */
}

func main2() (x int, y string) {
	abc = 1
	y = "abc"
	return
}
