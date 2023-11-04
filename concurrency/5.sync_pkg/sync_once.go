package main

import (
	"fmt"
	"sync"
)

var (
	only_once sync.Once
)

func sync_once() {
	for i := 0; i < 4; i++ {
		check()
	}
}

func check() {
	only_once.Do(func() {
		fmt.Println("printed only 1 time ")
	})
	fmt.Println("in check ")
}
