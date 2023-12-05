package main

import (
	"fmt"
	"math"
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

	divid, err3 := uint128.FromString(dividend.F)
	if err3 != nil {
		err = err3
		return
	}
	if divid.IsZero() {
		r = y.String()
		return
	}
	a, b, err2 := custom_div(divid, divisor, maxDecimalCount, roundingMode)
	if err2 != nil {
		err = err2
		return
	}

	f := div_by_10(a.String(), b.String(), len(dividend.F), maxDecimalCount)

	if y.IsZero() {
		r = f
		return
	}
	if fzero, _ := uint128.FromString(f); fzero.IsZero() {
		r = y.String()
		return
	}

	if strings.HasPrefix(f, "0") {
		var ov bool
		r, ov = add_fraction(y.String(), f)
		if ov {
			q = q.Add64(1)
		}
	} else {
		fint, _ := uint128.FromString(f)
		fractionSum := y.Add(fint)
		if OverFlow(fractionSum, y, f) {
			r = fractionSum.String()[1:]
			q = q.Add64(1)
		}
	}
	return
}

func custom_div(dividend, divisor uint128.Uint128, maxDecimalCount, roundingMode int) (q, r uint128.Uint128, err error) {
	if divisor.IsZero() {
		err = fmt.Errorf("division by zero")
		return
	}

	q, remainder := dividend.QuoRem(divisor)

	dividend = remainder.Mul64(uint64(math.Pow10(maxDecimalCount)))
	var roundOff uint128.Uint128
	remainder, roundOff = dividend.QuoRem(divisor)

	switch roundingMode {
	case 0: // to nearest
		lastDecimal := roundOff.Mul64(10).Div(divisor)
		if lastDecimal.Cmp64(5) > 0 {
			remainder = remainder.Add64(1)
		}
	case 1: // to up
		lastDecimal := roundOff.Mul64(10).Div(divisor)
		if lastDecimal.Cmp64(5) >= 0 {
			remainder = remainder.Add64(1)
		}
	case 2: // to down

	}
	r = remainder

	return
}
