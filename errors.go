package main

import (
	// "errors"
	"fmt"
)

type DivisionError struct{
	errorMessage string
}

func (e *DivisionError) Error() string {
	return e.errorMessage
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, &DivisionError{errorMessage: "Divide operation can not be done for having denominator as 0"}
		// return 0, errors.New("Divide operation can not be done for having denominator as 0")
	}
	return a / b, nil
}

func err() {
	result, err := divide(10, 0)
	if err == nil {
		fmt.Println("Result:", result)
	} else {
		fmt.Println("Error:", err)
	}
}
