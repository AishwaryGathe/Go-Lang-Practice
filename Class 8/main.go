package main

import (
	"fmt"
)

func main(){
	fmt.Println("Welcome to Array")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "PinApple" 
	fruitList[2] = "Manago" 
	fruitList[3] = "Banana" 

	fmt.Println("Fruit List is :",fruitList)
	fmt.Println("Fruit List is :", len(fruitList))

	var vegList = [3]string{"Potato", "moth", "Chana"}
	fmt.Println("Vegy List Is:", vegList)
	fmt.Println("Vegy List is:", len(fruitList))

}