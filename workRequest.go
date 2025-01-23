package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Course struct {
	CourseId   int     `json: "id"`
	CourseName string  `json: "name"`
	Price      float64 `json: "price"`
	Instructor string  `json: "instructor"`
}

var CourseList []Course

func init() {
	CourseJSON := `[
		{
			"id":1,
			"name":"Python",
			"price":2550,
			"instructor":"Frederick"
		},
		{
			"id":2,
			"name":"SQL",
			"price":3000,
			"instructor":"Andrew"
		},
		{
			"id":1,
			"name":"JavaScript",
			"price":1550,
			"instructor":"Frederick"
		}
	]`
	err := json.Unmarshal([]byte(CourseJSON), &CourseList)
	if err != nil {
		log.Fatal(err)
	}
}

func courseHandler(w http.ResponseWriter, r *http.Request) {
	courseJSON, err := json.Marshal(CourseList)
	
}

func main() {

}