package main

import (
	"encoding/json"
	"fmt"
)

type course struct{
	Name string `json:"coursename"`
	Price int	
	Platform string `json:"website"`
	Password string `json:"-"`
	Tags []string	`json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to Json advance")
	encodeJson()
}


func encodeJson(){
	lcoCourse := []course{
		{"ReactJS Bootcamp", 399, "aishwarygathe.tech", "abc@123", []string{"Web-dev","js"}},
		{"MERN Bootcamp", 199, "aishwarygathe.tech", "abc@321", []string{"Full-Stack","js"}},
		{"AngularJS Bootcamp", 299, "aishwarygathe.tech","abc@213", nil},
	}

	//package a data

	finalJson, err := json.MarshalIndent(lcoCourse, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}