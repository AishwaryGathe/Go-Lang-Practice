package main

import (
	"fmt"
)

func main(){
	fmt.Println("welcome to maps!")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["py"] = "python"
	languages["RB"] = "Ruby"
	languages["Cpp"] ="CPP"

	// fmt.Println("List of all Language: ", languages)
	fmt.Println("JS sort for : ", languages["JS"])
	
	delete(languages,"RB")
	fmt.Println("Lang: ", languages)


	// Loops are interesting in golang

	for _, value := range languages{
		fmt.Printf("For values is %v\n",value)
	}

}