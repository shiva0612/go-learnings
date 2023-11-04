package main

import (
	"fmt"
	"time"
)

func labelBreak() {

loop:
	for {
		select {
		case <-time.After(2 * time.Second):
			fmt.Println("test")
		case <-time.After(1 * time.Second):
			break loop
		}
	}
}
