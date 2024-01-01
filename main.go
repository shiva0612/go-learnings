package main

import (
	"fmt"
	"strings"

	"lukechampine.com/uint128"
)

type Decimal struct {
	V uint128.Uint128
	F string
}

func main() {
	run_example()
}

/*
123.0/1_00_000 = 0.00123 => but giving 0.1
123.123/3 = 41.041 => but giving 41.41
*/
func run_example() {

	dividend, divisor := get_input("1000000000000000000000", "2")
	maxDecimalCount := 3

	quotient, remainder, err := run_division(dividend, divisor, maxDecimalCount, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("%s / %s => %s, %s, nil\n", dividend, divisor, quotient, remainder)
	}

}

/*
123.456/3
123/3 + 456/3000
X.YY    A.BB
2.21	0.0156
q = X+A
r = YY+BB

	21
	0156
	----
	2256  (also if this len(ans) exceeds maxlen(a,b) => carry => q+1) -> DONE
	----
*/
func run_division(dividend Decimal, divisor uint128.Uint128, maxDecimalCount, roundingMode int) (q uint128.Uint128, r string, err error) {

	defer func() {
		if len(r) != maxDecimalCount {
			r += strings.Repeat("0", maxDecimalCount-len(r))
		}
	}()

	x, y, err1 := custom_div(dividend.V, divisor, maxDecimalCount, roundingMode)
	if err1 != nil {
		err = err1
		return
	}
	q = x

	divid, err3 := uint_from_string(dividend.F)
	if err3 != nil {
		err = err3
		return
	}
	if divid.IsZero() {
		r = y
		return
	}
	a, b, err2 := custom_div(divid, divisor, maxDecimalCount, roundingMode)
	if err2 != nil {
		err = err2
		return
	}

	f := div_by_10(a.String(), b, len(dividend.F), maxDecimalCount)

	if yzero, _ := uint_from_string(y); yzero.IsZero() {
		r = f
		return
	}
	if fzero, _ := uint_from_string(f); fzero.IsZero() {
		r = y
		return
	}

	if strings.HasPrefix(f, "0") {
		var ov bool
		r, ov = add_fraction(y, f)
		if ov {
			q = q.Add64(1)
		}
	} else {
		// fint, _ := uint_from_string(f)
		// fractionSum := y.Add(fint)
		// if OverFlow(fractionSum, y, f) {
		// 	r = fractionSum.String()[1:]
		// 	q = q.Add64(1)
		// }
	}
	return
}

func pre_zeros(b string) int {
	return len(b) - len(strings.TrimLeft(b, "0"))
}

/*
DONE y=12,f=012
DONE y=012,f=12

DONE y=12,f=12

DONE y=012,f=012
y=012,f=0012
y=0012,f=012
*/
func add_fraction2(a, b string) (string, bool) {

	az, bz := pre_zeros(a), pre_zeros(b)
	if az+bz == 0 {
		aint, _ := uint_from_string(a)
		bint, _ := uint_from_string(b)
		sum := aint.Add(bint)
		if OverFlow(sum, aint, bint.String()) {
			return sum.String()[1:], true
		}
		return sum.String(), false
	}
	no_of_zeros := bz

	if az != 0 && bz == 0 {
		a, b = b, a
		no_of_zeros = az

	}
	if az != 0 && bz != 0 {
		if az == bz {
			a, b = matchLength(a, b)
			aint, _ := uint_from_string(a)
			bint, _ := uint_from_string(b)
			sum := aint.Add(bint)
			return strings.Repeat("0", len(a)-len(sum.String())) + sum.String(), false

		} else if az > bz {

		} else {

		}
	}

	//b = 0125 (having prefixed with zeros)

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

func custom_div(dividend, divisor uint128.Uint128, maxDecimalCount, roundingMode int) (q uint128.Uint128, r string, err error) {
	if divisor.IsZero() {
		err = fmt.Errorf("division by zero")
		return
	}

	q, remainder := dividend.QuoRem(divisor)

	if remainder.IsZero() {
		r = strings.Repeat("0", maxDecimalCount)
		return
	}
	remStr := div_remainder(remainder, divisor, maxDecimalCount)

	newRem := ""
	var rounding bool
	switch roundingMode {
	case 0: // to nearest >5 => +1
		if remStr[len(remStr)-1] > '5' {
			rounding = true
		}
	case 1: // to up >=5 => +1
		if remStr[len(remStr)-1] >= '5' {
			rounding = true
		}
	case 2: // to down
		newRem = remStr[:maxDecimalCount]
	}
	if rounding {
		remint, _ := uint_from_string(remStr[:maxDecimalCount])
		remint = remint.Add64(1)
		if len(remint.String()) > maxDecimalCount {
			newRem = remint.String()[:maxDecimalCount]
			q = q.Add64(1)
		}
		newRem = remint.String()
	}

	if maxDecimalCount-len(newRem) > 0 {
		r = newRem + strings.Repeat("0", maxDecimalCount-len(newRem))
		return
	}
	r = newRem
	return
}

func div_remainder(divid, divis uint128.Uint128, maxdDecCount int) string {

	ans := ""
	for !divid.IsZero() && len(ans) <= maxdDecCount {
		if divid.Cmp(divis) < 0 {
			var times int
			divid, times = div_mul_10(divid, divis)
			ans += strings.Repeat("0", times)
		}
		q, r := divid.QuoRem(divis)
		ans += q.String()
		divid = r
	}
	if len(ans) < maxdDecCount+1 {
		ans += strings.Repeat("0", maxdDecCount+1-len(ans))
	}
	return ans[:maxdDecCount+1]
}

func div_mul_10(divid, divis uint128.Uint128) (uint128.Uint128, int) {
	times := 0
	divid = divid.Mul64(10)
	for divid.Cmp(divis) < 0 {
		divid = divid.Mul64(10)
		times += 1
	}
	return divid, times
}
