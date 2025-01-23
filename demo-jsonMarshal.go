package main

import (
	"encoding/json"
	"fmt"
)

type employee struct {
	ID           int
	EmployeeName string
	Tel          string
	Email        string
}

func main() {
	data, _ := json.Marshal(&employee{101, "Rodney Williams", "0900000000", "rodney@email.com"})
	fmt.Println(string(data))
}