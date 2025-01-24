package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Course struct {
	Id   int     `json: "id"`
	Name string  `json: "name"`
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
	switch r.Method {
	case http.MethodGet:
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(courseJSON)
	
	case http.MethodPost:
		var newCourse Course
		Bodybyte, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader((http.StatusBadRequest))
			return
		}
		err = json.Unmarshal(Bodybyte, newCourse)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		CourseList = append(CourseList, newCourse)
		w.WriteHeader(http.StatusCreated)
		
	}
	
}

func main() {
	http.HandleFunc("/course", courseHandler)
	http.ListenAndServe(":5000", nil)

}