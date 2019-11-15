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
		Data{10, 20},
		Data{30, 40},
		Data{50, 60},
		Data{70, 80},
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
