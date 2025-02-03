package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Course struct {
    Id         int     `json:"id"`
    Name       string  `json:"name"`
    Price      float64 `json:"price"`
    Instructor string  `json:"instructor"`
}

var CourseList []Course

func init() {
    CourseJSON := `[{"id":101,"name":"Python","price":2550,"instructor":"Frederick"},
                    {"id":102,"name":"SQL","price":3000,"instructor":"Andrew"},
                    {"id":103,"name":"JavaScript","price":1550,"instructor":"Frederick"}]`
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

func findID(ID int) (*Course, int) {
    for i, course := range CourseList {
        if course.Id == ID {
            return &CourseList[i], i
        }
    }
    return nil, -1
}

func courseHandler(w http.ResponseWriter, r *http.Request) {
    urlPathSegment := strings.Split(r.URL.Path, "course/")
    ID, err := strconv.Atoi(urlPathSegment[len(urlPathSegment)-1])
    if err != nil {
        log.Print(err)
        w.WriteHeader(http.StatusNotFound)
        return
    }

    course, listItemIndex := findID(ID)
    if listItemIndex == -1 {
        http.Error(w, fmt.Sprintf("No course found with ID %d", ID), http.StatusNotFound)
        return
    }

    switch r.Method {
    case http.MethodGet:
        json.NewEncoder(w).Encode(course)

    case http.MethodPatch:
        var updatedCourse Course
        byteBody, err := io.ReadAll(r.Body)
        if err != nil {
            w.WriteHeader(http.StatusBadRequest)
            return
        }

        err = json.Unmarshal(byteBody, &updatedCourse)
        if err != nil {
            w.WriteHeader(http.StatusBadRequest)
            return
        }

        updatedCourse.Id = course.Id

		//--------------------------//
		if updatedCourse.Name != "" {
            course.Name = updatedCourse.Name
        }
		if updatedCourse.Price != 0 {
            course.Price = updatedCourse.Price
        }
		if updatedCourse.Instructor != "" {
            course.Instructor = updatedCourse.Instructor
        }

		//-------------------------//

        CourseList[listItemIndex] = updatedCourse
        w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(CourseList)
    default:
        w.WriteHeader(http.StatusMethodNotAllowed)
    }
}

func coursesHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        json.NewEncoder(w).Encode(CourseList)

    case http.MethodPost:
        var newCourse Course
        Bodybyte, err := io.ReadAll(r.Body)
        if err != nil {
            w.WriteHeader(http.StatusBadRequest)
            return
        }

        err = json.Unmarshal(Bodybyte, &newCourse)
        if err != nil {
            w.WriteHeader(http.StatusBadRequest)
            return
        }

        if newCourse.Id != 0 {
            w.WriteHeader(http.StatusBadRequest)
            return
        }

        newCourse.Id = getNextId()
        CourseList = append(CourseList, newCourse)

        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(newCourse)
    }
}

func main() {
    http.HandleFunc("/course/", courseHandler)
    http.HandleFunc("/course", coursesHandler)
    http.ListenAndServe(":5000", nil)
}
