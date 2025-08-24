package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to loops in golang")

	days := []string{"sunday","monday","tuesday","wednesday","thursday","friday","saturday"}

	fmt.Println(days)

	// for d:=0; d<len(days);d++{
	// 	fmt.Println(days[d])
	// }

	// for i:=range days{
	// 	fmt.Println(days[i])
	// }

	// for _, day := range days{
	// 	fmt.Printf("Index is %v\n",day)
	// }

	rougueValue := 1

	for rougueValue <10{

		if rougueValue ==2{
			goto ajg
		}

		if rougueValue == 5{
			rougueValue++
			continue
		}

		fmt.Println("Value is :",rougueValue)
		rougueValue++
	}

	ajg:
		fmt.Println("Jumping at Aisharygathe")

}