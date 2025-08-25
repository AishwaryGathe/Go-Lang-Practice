package main

import (
	"fmt"
)

func main(){
	fmt.Println("Welcome to Methods!!!")

	// No inheritance in golang; No super or parent

	Aishwary := User{"Aishwary","aishwarygathe@gmail.com",true,16}
	fmt.Println(Aishwary)
	fmt.Printf("Aishwary Details are: %+v\n", Aishwary)
	fmt.Printf("Name is %v and email is %v\n", Aishwary.Name, Aishwary.Email)
	Aishwary.GetStatus()
	Aishwary.NewMail()
	fmt.Printf("Name is %v and email is %v\n", Aishwary.Name, Aishwary.Email)
}

type User struct{
	Name string
	Email string
	Status bool
	Age int
}

func (u User) GetStatus(){
	fmt.Println("Is user Active:",u.Status)
}

func (u User) NewMail(){
	u.Email = "aishwarygathe20@gmail.com"
	fmt.Println("Email of this user :",u.Email)
}