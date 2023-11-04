package main

import (
	"fmt"
	"sync"
)

type Person struct {
	Count      int
	sync.Mutex //mutex should always be non-pointer
	/*
		if you say *sync.Mutex
		then while creating p := Person{}, you have to exclusively say p.Mutex = &sync.Mutex{}
	*/
}

var (
	p  = &Person{}
	wg sync.WaitGroup
)

func main() {

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go incr()
	}
	wg.Wait()
	fmt.Println("count = ", p.Count)
}
func incr() {

	defer wg.Done()
	p.Lock()
	p.Count += 1
	p.Unlock()
}
