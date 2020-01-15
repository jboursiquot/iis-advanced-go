package main

import (
	"fmt"
	"testing"
)

func Setup(t *testing.T)    { t.Log("Unit test setup") }
func TearDown(t *testing.T) { t.Log("Unit test tear down") }

func Cleanup() { fmt.Println("suite cleanup...") }
func TestMain(m *testing.M) {
	// do some setup, parse flags, etc
	fmt.Println("suite setup...")
	defer Cleanup()
	m.Run()
	return
}

func TestDeposit(t *testing.T) {
	Setup(t)
	defer TearDown(t)

	balance := float32(100)
	deposit := float32(200)
	if Deposit(balance, deposit) != 300 {
		t.Error("Unit test for Deposit failed")
	}
}
func TestWithdrawal(t *testing.T) {
	balance := float32(400)
	withdrawal := float32(200)
	result, _ := Withdrawal(balance, withdrawal)
	if result != 200 {
		t.Error("Unit test for Withdrawal failed")
	}
}
