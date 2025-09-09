package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//model for course -file
type Course struct{
	CourseId string `json:"courseid"`
	CourseName string `json:"coursename"`
	CoursePrice int  `json:"price"`
	Author *Author `json:"author"`	
}

type Author struct{
	Fullname string `json:"fullname"`
	Website string `json:"website"`
}



//fake Database
var courses []Course

//middleware -file

func (c *Course) Isempty() bool {
	return c.CourseId == "" && c.CourseName == ""
}

func main() {
	fmt.Println("Hello API!!!")

}

//controllers - file
// serve home route
func serveHome(w http.ResponseWriter , r *http.Request){
	w.Write([]byte("<h1>Welcome to APi by aishwary</h1>"))
}

func getAllCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Get all Courses")
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOnecourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Get one course")
	w.Header().Set("Content-Type","application/json")

	params := mux.Vars(r)
	fmt.Println(params)

	for _, course := range courses{
		if course.CourseId == params["id"]{
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course foudn with given id")
}



