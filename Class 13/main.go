package main

import(
	"fmt"
)

func main(){
	fmt.Println("Welcome to IF ELSE!!!")
	// loginCount := 10

	// var result string

	// if loginCount < 10 {
	// 	result ="Regular User"
	// } else if loginCount >10 {

	// 	result="Watch Out"
	// } else {
	// 	result="Exactly 10 login count"
	// }


	// fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Number is Even")
	} else {
		fmt.Println("Number is odd")
	}

	if num := 3; num <10{
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is NOT less than 10")
	}
}