package main

import (
	"fmt"
)

func main() {
	fmt.Println("welcome to class pointer")

	// var ptr *int
	// fmt.Println("Value of pointer is  ", ptr)

	myNumber := 23
	var ptr = &myNumber

	fmt.Println("Value of the Pointer ",ptr)
	fmt.Println("Value of the Pointer ",*ptr)

	*ptr = *ptr + 2
	fmt.Println("New Value is :", myNumber)
}