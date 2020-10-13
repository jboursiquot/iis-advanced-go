package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
)

// Actor is...
type Actor struct {
	Name   string `json:"Actor"`
	Role   string
	Salary int
}

func main() {
	const jsonStream = `
		{"Actor": "Robert", "Role": "Bob", "Salary": 500000}
		{"Actor": "Elliot", "Role": "Ted", "Salary": 250000}
        {"Actor": "Natalie", "Role": "Carol", "Salary": 450000}
		{"Actor": "Dyan", "Role": "Alice", "Salary": 200000}
	`

	dec := json.NewDecoder(strings.NewReader(jsonStream))
	for {
		var a Actor
		if err := dec.Decode(&a); err == io.EOF {
			break
		} else if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("%s, %s, %d\n", a.Name, a.Role, a.Salary)
	}
}
