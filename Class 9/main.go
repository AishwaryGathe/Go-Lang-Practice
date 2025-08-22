package main

import (
	"fmt"
	"sort"
)

func main(){
	fmt.Println("Welcome to Slices")
	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of FruitList is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "banana")

	fmt.Println(fruitList)

	// fruitList = append(fruitList[1:3])
	// fmt.Println(fruitList)

	// fruitList =append(fruitList[:3])
	// fmt.Println(fruitList)


	highscore := make([]int, 4)

	highscore[0] = 234
	highscore[1] = 934
	highscore[2] = 434
	highscore[3] = 534
	// highscore[4] = 777

	highscore = append(highscore, 555,66,777)

	fmt.Println(highscore)

	fmt.Println(sort.IntsAreSorted(highscore))
	sort.Ints((highscore))
	fmt.Println(highscore)


}