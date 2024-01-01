package main

import (
	"strings"

	"lukechampine.com/uint128"
)

func uint_from_string(a string) (uint128.Uint128, error) {
	a = strings.TrimLeft(a, "0")
	return uint128.FromString(a)
}

func getMaxLen(a, b string) int {
	lena := len(a)
	lenb := len(b)
	if lena > lenb {
		return lena
	}
	return lenb
}
func getMaxLenUint128(a, b uint128.Uint128) int {
	lena := len(a.String())
	lenb := len(b.String())
	if lena > lenb {
		return lena
	}
	return lenb
}

func OverFlow(sum, a uint128.Uint128, b string) bool {
	slen, alen, blen := len(sum.String()), len(a.String()), len(b)
	if alen > blen {
		if slen > alen {
			return true
		}
	} else {
		if slen > blen {
			return true
		}
	}
	return false
}

// 21
// 0156
// ----
// 2256
func add_fraction(a, b string) (string, bool) {

	//b = 0125 (having prefixed with zeros)
	no_of_zeros := len(b) - len(strings.TrimLeft(b, "0"))

	la := a[:no_of_zeros]
	laint, _ := uint_from_string(la)
	ra, rb := a[no_of_zeros:], b[no_of_zeros:]
	ra, rb = matchLength(ra, rb)
	raint, _ := uint_from_string(ra)
	rbint, _ := uint_from_string(rb)
	sumr := raint.Add(rbint)
	if OverFlow(sumr, raint, rbint.String()) {
		suml := laint.Add64(1)
		if OverFlow(suml, laint, "1") {
			return suml.String()[1:] + sumr.String()[1:], true
		}
		return suml.String() + sumr.String()[1:], false
	}
	return la + sumr.String(), false
}

/*
2.1		2.1000
0.0156  0.0156

0.014	0.01400
2.12355 2.12355
*/
func matchLength(a, b string) (string, string) {
	alen, blen := len(a), len(b)
	if alen > blen {
		return a, b + strings.Repeat("0", alen-blen)
	}
	return a + strings.Repeat("0", blen-alen), b
}

/*
41.333/1000 = 0.04133 = 04133
*/
func div_by_10(a, b string, den, maxDecCount int) string {
	f := strings.Repeat("0", den-len(a)) + a
	if len(f) == maxDecCount {
		return f
	}
	if len(f) > maxDecCount {
		return f[:maxDecCount]
	}
	return f + b[:maxDecCount-len(f)]
}
func removeFirstDigit(num uint64) int64 {
	var numDigits uint64 = 1
	for temp := num; temp >= 10; temp /= 10 {
		numDigits *= 10
	}

	result := num % numDigits
	return int64(result)
}

func get_input(dividend, divisor string) (Decimal, uint128.Uint128) {
	divid := Decimal{}
	var divis uint128.Uint128

	index := strings.IndexByte(dividend, '.')
	if index != -1 {
		divid.V, _ = uint_from_string(dividend[:index])
		divid.F = dividend[index+1:]
	} else {
		divid.V, _ = uint_from_string(dividend)
		divid.F = "0"
	}
	divis, _ = uint_from_string(divisor)
	return divid, divis
}
