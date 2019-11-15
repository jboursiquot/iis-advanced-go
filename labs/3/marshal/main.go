package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Employee is...
type Employee struct {
	Name      string
	Type      string
	Age       string
	Insurance string
}

func main() {
	data := []Employee{
		Employee{Name: "Bob Jones", Type: "Manager", Age: "55"},
		Employee{Name: "Sally Wilson", Type: "Vice President", Age: "45", Insurance: "No"},
		Employee{Name: "Sam Smith", Type: "Warehouse", Insurance: "Yes"},
		Employee{Name: "Adam Freeman", Type: "Warehouse", Age: "34", Insurance: "Yes"},
	}

	bs, err := json.Marshal(data)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(bs))
}
