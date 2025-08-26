package main

import (
	"fmt"
	"net/url"
)

const myurl string ="https://lco.dev:3000/learn?coursename=reactjs&paymentighbj456ghb"

func main() {
	fmt.Println("Welcome to handling URLs")

	fmt.Println(myurl)

	result, err := url.Parse(myurl)

	if err != nil {
		panic(err)
	}

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The type of query params are :%T\n",qparams)

	fmt.Println(qparams["coursename"])

	for _, val := range qparams{
		fmt.Println("Param is:",val)
	}

	partofUrl := &url.URL{
		Scheme: "https",
		Host: "loc.dev",
		Path: "/learn",
		RawPath: "user=hitesh",
	}

	anotherURL := partofUrl.String()
	fmt.Println(anotherURL)


}