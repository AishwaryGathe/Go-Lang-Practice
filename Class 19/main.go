package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files!!!")

	content := "This needs to go in file - aishwarygath.tech"

	file, err := os.Create("./myFile.txt")
	checkNilerr(err)

	length, err := io.WriteString(file, content)
	checkNilerr(err)
	fmt.Println("Length is:", length)
	defer file.Close()

	readFile("./myFile.txt")

}

func readFile(fileName string) {
	databyte, err := ioutil.ReadFile(fileName)
	checkNilerr(err)
	fmt.Println("Text data inside the file is :\n", string(databyte))
}

func checkNilerr(err error) {
	if err != nil {
		panic(err)
	}
}
