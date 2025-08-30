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
	// encodeJson()
	DecodeJson()
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


func DecodeJson(){
	jsonDataWeb := []byte(`
	{
                "coursename": "ReactJS Bootcamp",
                "Price": 399,
                "website": "aishwarygathe.tech", 
                "tags": ["Web-dev","js"]
    }
		`)

		var lcoCourse course

		checkValid := json.Valid(jsonDataWeb)

		if checkValid {
			fmt.Println("JSON was Valid!!")
			json.Unmarshal(jsonDataWeb,&lcoCourse)
			fmt.Printf("%#v\n", lcoCourse)
			}	else {
				fmt.Println("JSON was NOT VALID!!")
			}
			
			// some cases 
			
			var myOnlineData map[string]interface{}
			json.Unmarshal(jsonDataWeb,&myOnlineData)
			fmt.Printf("%#v\n", myOnlineData)


			for k,v := range myOnlineData{
				fmt.Printf("Key is %v and value is %v and type us %T\n", k, v,v)
			}
}