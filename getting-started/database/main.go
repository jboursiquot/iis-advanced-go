package main

import (
	"encoding/json"
	"fmt"

	scribble "github.com/nanobox-io/golang-scribble"
)

type Person struct{ Name string }

func main() {
	db, _ := scribble.New("./persons", nil)
	for _, name := range []string{"Bob", "Ted", "Carol", "Alice"} {
		db.Write("persons", name, Person{Name: name})
	} // add data

	p := Person{}
	db.Read("persons", "Bob", &p) // Read person
	fmt.Printf("Name is %#v\n", p)

	actors, _ := db.ReadAll("persons") // Read everyone
	persons := []Person{}
	for _, actor := range actors {
		json.Unmarshal([]byte(actor), &p)
		persons = append(persons, p)
	}
	fmt.Printf("Movie actors! %#v\n", persons)
	db.Delete("persons", "")
}
