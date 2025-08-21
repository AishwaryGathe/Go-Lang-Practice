package main

import (
	"fmt"
	"time"
)

func main(){
	fmt.Println("Welcome to time study of golang")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04 Monday"))

	createDate := time.Date(2025, time.August, 20, 23, 23, 0, 0,time.UTC )
	fmt.Println(createDate.Format("01-02-2006 Monday"))
}

// go env