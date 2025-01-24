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
			"id":101,
			"name":"Python",
			"price":2550,
			"instructor":"Frederick"
		},
		{
			"id":102,
			"name":"SQL",
			"price":3000,
			"instructor":"Andrew"
		},
		{
			"id":103,
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

func getNextId() int {
	highestId := -1
	for _, course := range CourseList {
		if highestId < course.Id {
			highestId = course.Id
		}
	}
	return highestId + 1
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
		err = json.Unmarshal(Bodybyte, &newCourse)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		//Create new id
		if newCourse.Id != 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		newCourse.Id = getNextId()


		CourseList = append(CourseList, newCourse)
		w.WriteHeader(http.StatusCreated)
		
	}
	
}

func main() {
	http.HandleFunc("/course", courseHandler)
	http.ListenAndServe(":5000", nil)

}