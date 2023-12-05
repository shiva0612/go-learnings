package main

import (
	"fmt"
	"sort"
)

func sorting() {
	a := [][]int{{1, 10, 3}, {3, 4, 5}, {5, 6, 7}}

	/*
		if u use "<"
			a<b<c<d - ascending
		if u use ">"
			a>b>c>d - descending
	*/
	sort.Slice(a, func(i, j int) bool {
		return a[i][1] < a[j][1]
	})

	fmt.Println(a)
}
