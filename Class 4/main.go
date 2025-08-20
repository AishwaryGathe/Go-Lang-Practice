package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


func main(){
	fmt.Println("Welcome to Our Pizza App")
	fmt.Println("Please Rate Our Pizza Between 1 and 5")
	
	rating := bufio.NewReader(os.Stdin)

	input, _ := rating.ReadString('\n')
	fmt.Println("Thanks for the Rating",input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input),64)

	if err != nil{
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your Rating: ", numRating+1)
	}

}