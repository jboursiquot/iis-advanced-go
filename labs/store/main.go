package main

import (
	"fmt"
	"sync"
)

type store struct {
	sync.RWMutex
	sales  int
	people int
}

func (s *store) shop(customer string, amount int) {
	s.Lock()
	defer s.Unlock()
	s.people++
	s.sales += amount
	fmt.Printf("%s spent %d\n", customer, amount)
}

func (s *store) ledger() {
	s.RLock()
	defer s.RUnlock()
	fmt.Printf("%d person(s) spent %d dollars in store\n", s.people, s.sales)
}

func main() {
	s := new(store)

	var wg sync.WaitGroup
	wg.Add(4)

	go func() { s.shop("Bob", 1000); wg.Done() }()
	go func() { s.shop("Ted", 2000); wg.Done() }()
	go func() { s.shop("Carol", 3000); wg.Done() }()
	go func() { s.shop("Alice", 4000); wg.Done() }()

	wg.Wait()

	lchan := make(chan struct{})

	go func(ch <-chan struct{}, wg *sync.WaitGroup) {
		<-ch
		s.ledger()
		wg.Done()
	}(lchan, &wg)

	wg.Add(1)
	lchan <- struct{}{}
	wg.Wait()
}
