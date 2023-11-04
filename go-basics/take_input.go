package main

import "fmt"

func take_input() {
	var i, j int
	fmt.Scan(&i)
	fmt.Scanln(&i, &j)
	fmt.Scanf("%v %v", &i, &j)
}
