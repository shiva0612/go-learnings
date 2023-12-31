package main

import (
	"fmt"
	"sync"
)

var x = 0

func increment(wg *sync.WaitGroup, ch chan bool) {
	ch <- true //another go routine cannot write into this since buffered channel of length = 1
	x = x + 1
	<-ch
	wg.Done()
}
func main() {
	var w sync.WaitGroup
	/*
		it cannot be un-buffered channel bcz all the go-routines are writing into channel and there is no go-routine which is reading
		so, dead lock
	*/
	ch := make(chan bool, 1) //buffered channel of len=1 (IMP)
	defer close(ch)
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w, ch)
	}
	w.Wait()
	fmt.Println("final value of x", x)
}
