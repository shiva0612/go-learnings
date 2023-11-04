package main

import "fmt"

func interfaces() {
	// Define an interface{} variable with an underlying []string
	var data interface{}
	data = []string{"apple", "banana", "cherry"}

	// Type assert data to a []string
	if strings, ok := data.([]string); ok {
		// Now you can work with strings as a []string
		for _, s := range strings {
			fmt.Println(s)
		}
	} else {
		fmt.Println("data is not a []string")
	}
}
