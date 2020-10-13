package main

import (
	"encoding/json" //JSON package from the standard library
	"fmt"
	"os"
)

type Data struct {
	First  int
	Second int
}

func main() {
	var file, _ = os.Create("data.txt")
	data := []Data{
		{10, 20},
		{30, 40},
		{50, 60},
		{70, 80},
	}
	count := len(data)
	enc := json.NewEncoder(file)
	file.Write([]byte("["))

	for index, record := range data {
		if err := enc.Encode(&record); err != nil {
			fmt.Println(err.Error())
			return
		}
		if (count - 1) == index {
			continue
		}
		file.Write([]byte(","))
	}
	file.Write([]byte("]"))
	file.Close()
}
