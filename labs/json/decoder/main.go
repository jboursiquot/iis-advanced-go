package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Data struct {
	First  int
	Second int
}

func main() {
	var file, _ = os.Open("data.txt")
	dec := json.NewDecoder(file)
	dec.Token()
	for dec.More() {
		var data Data
		dec.Decode(&data)
		fmt.Println(data.First, data.Second)
	}
	file.Close()
}
