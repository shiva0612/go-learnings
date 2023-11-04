package main

import (
	"fmt"
	"log"
)

func main() {
	// enum_test()
	// replacing_in_strings()
	// take_input()
	// labelBreak()
	// wrapping_error()
	// representation_bytes()
	sorting()
	// testing()
}

//------------------------------------------------------------------
/*
2023/11/03 22:34:29 MAIN: list = 0x140000160e0
2023/11/03 22:34:29 func: listp = 0x1400000c048
2023/11/03 22:34:29 list =  [1 2 3]
2023/11/03 22:34:29 func: listp = 0x1400000c048
[1 2]
*/
func testing() {
	a := []int{1, 2}
	log.Printf("MAIN: list = %p", a)
	f(a)
	fmt.Println(a)
}

func f(list []int) {
	listp := &list
	log.Printf("func: listp = %p", listp)

	*listp = append(*listp, 3)
	log.Println("list = ", list)

	log.Printf("func: listp = %p", listp)

}

//------------------------------------------------------------------
