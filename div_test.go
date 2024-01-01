package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	dividend string
	divisor  string
	q        string
	r        string
}

func TestDiv(t *testing.T) {

	cases := []TestCase{
		// {"123.123", "3", "41", "041"},
		// {"987654321987654321.987654321987654321", "123456789123456789", "8", "000"},
		// {"1000000000000000000000", "2", "500000000000000000000", "000"},
		// {"0", "123123123", "0", "000"},
		// {"123123123.123123123", "1", "123123123", "123"},
		// {"83703.168", "678", "123", "456"},
		{"38612.9032", "30", "1287", "096"},
		// {"36.1", "4", "9", "025"},
	}

	for _, test_ip := range cases {
		t.Run(fmt.Sprintf("%s/%s=%s.%s", test_ip.dividend, test_ip.divisor, test_ip.q, test_ip.r), func(t *testing.T) {
			divid, divis := get_input(test_ip.dividend, test_ip.divisor)
			q, r, err := run_division(divid, divis, 3, 2)
			if err != nil {
				t.Logf("failed with error : %s", err.Error())
				t.Fail()
			}

			if q.String() != test_ip.q {
				t.Logf("expected(q) = %s, got(q) = %s", test_ip.q, q)
				t.Fail()
			}
			if r != test_ip.r {
				t.Logf("expected(r) = %s, got(r) = %s", test_ip.r, r)
				t.Fail()
			}
		})
	}
}

func TestAddFraction(t *testing.T) {

	type TC struct {
		first, second, ans string
		overflow           bool
	}

	cases := []TC{
		{
			"29", "0156", "3056", false,
		},
		{
			"99", "0156", "0056", true,
		},
		{
			"12", "12", "24", false,
		},
		{
			"91", "91", "82", true,
		},
		{"21", "0156", "2256", false},
		{"012", "12", "132", false},
		{"12", "012", "132", false},
		{"09", "01", "10", false},
		{"009", "001", "010", false},
		{"002", "003", "005", false},
		{"001", "0012", "0022", false},
	}

	for _, test_case := range cases {
		t.Run(fmt.Sprintf("%s+%s", test_case.first, test_case.second), func(t *testing.T) {
			got, overflow := add_fraction2(test_case.first, test_case.second)
			assert.Equal(t, test_case.ans, got)
			assert.Equal(t, test_case.overflow, overflow)
		})
	}
}

func TestMatchLength(t *testing.T) {
	// 2.1		2.1000
	// 0.0156  0.0156

	// 0.014	0.01400
	// 2.12355 2.12355

	a, b := "14", "2355"
	gota, gotb := matchLength(a, b)
	assert.Equal(t, "1400", gota)
	assert.Equal(t, "2355", gotb)

}

func TestCustomDiv(t *testing.T) {

	cases := []TestCase{
		{"38612", "30", "1287", "066"},
		{"9032", "30", "301", "066"},
	}

	for _, test_ip := range cases {
		t.Run(fmt.Sprintf("%s/%s=%s.%s", test_ip.dividend, test_ip.divisor, test_ip.q, test_ip.r), func(t *testing.T) {
			divid, divis := get_input(test_ip.dividend, test_ip.divisor)
			q, r, err := custom_div(divid.V, divis, 3, 2)
			if err != nil {
				t.Fail()
				t.Log(err.Error())
			}
			assert.Equal(t, test_ip.q, q.String())
			assert.Equal(t, test_ip.r, r)
		})
	}
}
