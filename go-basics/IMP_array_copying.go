package main

import "fmt"

func copying_slices() {
	// a := []int{0, 1, 2, 3, 4}
	// b := a[:2]
	// b[0] = 1
	// fmt.Println(a, b) //[1 1 2 3 4] [1 1]

	// a := []int{0, 1, 2, 3, 4}
	// b := []int{}
	// copy(b, a) // len of b is 0 => nothing is copied
	// b[0] = 1   //panic
	// fmt.Println(a, b)

	// a := []int{0, 1, 2, 3, 4}
	// b := make([]int, len(a[:2]))
	// copy(b, a[:2])
	// b[0] = 1
	// fmt.Println(a, b) //[0 1 2 3 4] [1 1]

	// a := []int{0, 1, 2, 3, 4}
	// b := append([]int{}, a[:2]...)
	// b[0] = 1
	// fmt.Println(a, b) //[0 1 2 3 4] [1 1]

	//if you want to copy some part of slice into another by value use copy - but make sure that the len of dest and src are same in copy
	a := []int{0, 1, 2, 3, 4}
	b := [3]int{}
	copy(b[:2], a[:2])
	b[0] = 1
	fmt.Println(a, b) //[0 1 2 3 4] [1 1 0]

}

func copying_2d_arrays() {

	// a := []int{1}
	// b := []int{1}
	// a = append(a, b...) //a = 1,1
	// b[0] = 2            //b = 2
	// fmt.Println(a, b) //a = 1,1   b = 2

	// a := &[]int{1}
	// b := &[]int{1}
	// *a = append(*a, *b...)
	// (*b)[0] = 2
	// fmt.Println(a, b) //same as above
	//----------------since we spread b it is like []int{b...}, so connection with b is lost while appending-------

	// a := [][]int{{1}} // a = [[1]]
	// b := []int{1}
	// a = append(a, b) // a = [[1] [1]]
	// b[0] = 2
	// fmt.Println(a, b) //[[1] [2]] since we are appending b which is slice which is nothing but pointer

	// a := [][]int{{1}} // a = [[1]]
	// b := []int{1}
	// a = append(a, append([]int{}, b...)) // a = [[1] [1]]
	// b[0] = 2
	// fmt.Println(a, b) // a = [[1] [1]]   b = [2]

	a := &[][]int{{1}}
	b := &[]int{1}
	*a = append(*a, append([]int{}, *b...))
	(*b)[0] = 2
	fmt.Println(a, b) //a = &[[1] [1]]   b = &[2]
}
