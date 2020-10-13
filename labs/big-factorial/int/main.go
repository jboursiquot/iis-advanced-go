package main

import "os"

import "strconv"

import "log"

func doTheWork(ch chan int, input int) {
	ch <- factorial(input)
}

func factorial(n int) int {
	var result int
	if n > 0 {
		result = n * factorial(n-1)
		return result
	}
	return 1
}

func main() {
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}

	ch := make(chan int)
	go doTheWork(ch, n)
	println(<-ch)
}
