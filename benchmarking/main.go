package main

import (
	"fmt"
	"strings"
)

func main() {

}

func concatnormal(a, b string) string {
	return a + b
}

func concatfmt(a, b string) string {
	return fmt.Sprintf("%s%s", a, b)
}

func concatstrings(a, b string) string {
	sb := strings.Builder{}
	sb.WriteString(a)
	sb.WriteString(b)
	return sb.String()
}

func benchStatTest(a, b string) string {
	// return concatfmt(a, b)
	return concatstrings(a, b)
}
