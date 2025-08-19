package main

import (
	"fmt"
)

const Logintoken string ="goodbye"


func main(){
	fmt.Println("Varibles")
	var username string = "Aishwary"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n",username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Varibles is of type: %T \n",isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Varibles is of Type: %T \n",smallVal)

	var smallFloat float64 = 255.4544444444444
	fmt.Println(smallFloat)
	fmt.Printf("Vraiables is of type : %T \n", smallFloat)

	var anothervar int
	fmt.Println(anothervar)
	fmt.Printf("Vriable is of type: %T \n", anothervar)

	var website ="aishwarygathe.tech"
	fmt.Println(website)

	numberofuser := 3000.0
	fmt.Println(numberofuser)

	fmt.Println(Logintoken)
	fmt.Printf("Vriable is of type: %T \n", Logintoken)
}