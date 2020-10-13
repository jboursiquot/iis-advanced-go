package main

import (
	"errors"
	"fmt"
)

func Deposit(balance float32, amount float32) float32 {
	return balance + amount
}

func Withdrawal(balance float32, amount float32) (float32, error) {
	current := balance - amount
	if current < 0 {
		return balance, errors.New("Not allowed to overdraw account")
	} else {
		return current, nil
	}
}

func main() {
	balance := float32(1000)
	balance = Deposit(balance, 500)
	var err error = nil
	balance, err = Withdrawal(balance, 2000)
	if err == nil {
		fmt.Println("balance is ", balance)
	} else {
		fmt.Println(err.Error())
	}
}
