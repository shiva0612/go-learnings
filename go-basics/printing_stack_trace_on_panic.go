package main

import (
	"fmt"
	"runtime"
)

func stack_trace() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic occurred:", r)
			stack := make([]byte, 4096)
			length := runtime.Stack(stack, true)
			fmt.Printf("Stack Trace:\n%s\n", stack[:length])
		}
	}()

	// Some code that may potentially panic
	panic("Something went wrong")
}
