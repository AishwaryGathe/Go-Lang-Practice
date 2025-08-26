package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://loc.dev"

func main() {
	fmt.Println("Welcome to Web Request!!")

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Resposne Type: %T\n", response)
	defer response.Body.Close()

	databytes,err := ioutil.ReadAll(response.Body)
	if err!=nil{
		panic(err)
	}
	content := string(databytes)

	fmt.Println("Values:",content)
}
