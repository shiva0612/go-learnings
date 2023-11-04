package main

import "fmt"

/*
1bit
1byte = 8bits

1024bits = 1kbits --> internets speed 1kbps = 1kilo bits /sec = 1024 bits/sec => 1024/8=128bytes/sec
1024*1024bits = 1mbits
1024*1024*1024bits = 1gbits
1024*1024*1024*1024bits = 1tbits

1024bytes = 1kBytes
1024*1024bytes = 1mBytes
1024*1024*1024bytes = 1gBytes
1024*1024*1024*1024bytes = 1tBytes

from quora ...
1byte=8bits..
1 KB=kilo byte=1024bytes=8196 bits.
1 Kb=kilo bits=1024bits/8 bits=128 bytes.
*/

func representation_bytes() {
	a := 1 << 10       //1kb
	b := 100 * 1 << 10 //100kb
	c := 1 << 20       //1mb
	d := 100 * 1 << 20 //100mb
	e := 1 << 30       //1gb

	fmt.Println(a, b, c, d, e)
}

func bits_bytes() {
	a := 1024
	b := 1 << 10        //int(1 00000 00000) = 1024 is stored as int
	fmt.Println(a == b) //true
}

func bitwise_op() {
	var a uint = 60 /* 60 = 0011 1100 */
	var b uint = 13 /* 13 = 0000 1101 */
	var c uint = 0

	c = a & b /* 12 = 0000 1100 */
	fmt.Printf("Line 1 - Value of c is %d\n", c)

	c = a | b /* 61 = 0011 1101 */
	fmt.Printf("Line 2 - Value of c is %d\n", c)

	c = a ^ b /* 49 = 0011 0001 */
	fmt.Printf("Line 3 - Value of c is %d\n", c)

	c = a << 2 /* 240 = 1111 0000 */
	fmt.Printf("Line 4 - Value of c is %d\n", c)

	c = a >> 2 /* 15 = 0000 1111 */
	fmt.Printf("Line 5 - Value of c is %d\n", c)
}
