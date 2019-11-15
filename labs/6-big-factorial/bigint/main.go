package main

import (
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
)

func doTheWork(ch chan *big.Int, input int64) {
	ch <- factorial(input)
}

func factorial(n int64) (result *big.Int) {
	result = big.NewInt(1)
	if n > 0 {
		return result.Mul(big.NewInt(n), factorial(n-1))
	}
	return
}

func main() {
	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}

	ch := make(chan *big.Int)
	go doTheWork(ch, int64(n))
	fmt.Println(<-ch)
}
