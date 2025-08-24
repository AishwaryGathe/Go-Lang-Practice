package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	fmt.Println("Welcome to switchcase!!")

	rand.Seed(time.Now().UnixNano())
	diceNum := rand.Intn(6) + 1
	fmt.Println("Value of dice is:", diceNum)

	switch diceNum {
	case 1:
		fmt.Println("Dice Value is 1 and you can open")
	case 2:
		fmt.Println("You got 2!")
		fallthrough
	case 3:
		fmt.Println("You got 3!")
	case 4:
		fmt.Println("You got 4!")
	case 5:
		fmt.Println("You got 5!")
	case 6:
		fmt.Println("You got one more chance and 6 to!!")
	default:
		fmt.Println("NO SPOT What was that!")
	}
}