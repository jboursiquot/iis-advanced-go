package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	params := os.Args[1:]
	if len(params) != 3 {
		log.Fatalln("expected invocation with <filename> <old> <new>")
	}

	path := params[0]

	bs, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}

	newString := strings.Replace(string(bs), params[1], params[2], -1)

	if err := ioutil.WriteFile(path, []byte(newString), 0644); err != nil {
		log.Fatalf("error: %s\n", err)
	}
}
