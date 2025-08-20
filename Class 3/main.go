package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	welcome := "Welcome User to Code"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Pizza:")

	// Comma Ok || err ok

	input, _ := reader.ReadString('\n')
	fmt.Printf("Thanks for rating %T",input)

}