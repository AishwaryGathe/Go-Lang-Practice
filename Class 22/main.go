package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Hellow this is server!!")
	// PerformGetrequest()
	// PerformPostjsonRequest()
	PerformPostRequest()
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

	fmt.Println("ByteContent is:", byteCount)
	fmt.Println(responseString.String())

	// fmt.Println(string(content))

}

func PerformPostjsonRequest() {
	const myurl = "http://localhost:8000/post"

	//fake json requrst

	requestBody := strings.NewReader(`
		{
			"coursename":"lets go with golang....",
			"price":0,
			"platform":"aishwarygathe.tech"
		}
	`)

	response, err :=http.Post(myurl,"application/json",requestBody)

	if err != nil{
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}


func PerformPostRequest(){
	const myurl = "http://localhost:8000/postform"

	data := url.Values{}
	data.Add("firstname","Aishwary")
	data.Add("lastname","gathe")
	data.Add("email","Aishwarygathe@gmail.com")

	response, err := http.PostForm(myurl,data)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}