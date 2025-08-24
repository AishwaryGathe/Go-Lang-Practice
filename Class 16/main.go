package main

import (
	"fmt"
)

func main(){
	fmt.Println("Welcome to Functions")
	greeter()

	result := adder(3,5)
	fmt.Println("Result is :",result)

	proresult, mymessage := proAdder(2,2,2,2,2)
	fmt.Println("ProResults is :",proresult)
	fmt.Println("ProMessage is :",mymessage)
}

func adder(valueOne int, valueTwo int) int{
	return valueOne + valueTwo
}

func proAdder(values ...int) (int, string){
	total:= 0
	for _, val := range values{
		total +=val
	}

	return total, "Hi i,m proAdder!!"
}

func greeter(){
	fmt.Println("NAMASTE from AISHWARY!!!")
}