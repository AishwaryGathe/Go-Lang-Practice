package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Hellow this is server!!")
	PerformGetrequest()
}

func PerformGetrequest() {
	const myurl = "http://localhost:8000/get"
	response, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)

	fmt.Println("Content length is:", response.ContentLength)


	var responseString strings.Builder
	content, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	byteCount, _ := responseString.Write(content)

	fmt.Println("ByteContent is:",byteCount)
	fmt.Println(responseString.String())
	

	// fmt.Println(string(content))

}
