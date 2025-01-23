package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type employee struct {
	ID           int
	EmployeeName string
	Tel          string
	Email        string
}

func main() {

	myEmployee := employee{}
	err := json.Unmarshal([]byte(`{"ID":1, "EmployeeName":"Peter", "Tel":"0900000000", "Email":"peter@mail.com"}`), &myEmployee)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(myEmployee.EmployeeName)
}