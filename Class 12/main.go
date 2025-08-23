package main

import (
	"fmt"
)

func main(){
	fmt.Println("Welcome to Struct")

	// No inheritance in golang; No super or parent

	Aishwary := User{"Aishwary","aishwarygathe@gmail.com",true,16}
	fmt.Println(Aishwary)
	fmt.Printf("Aishwary Details are: %+v\n", Aishwary)
	fmt.Printf("Name is %v and email is %v\n", Aishwary.Name, Aishwary.Email)

	
}

type User struct{
	Name string
	Email string
	Status bool
	Age int
}