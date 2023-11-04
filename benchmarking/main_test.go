package main

import (
	"testing"
)

func BenchmarkNormal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ans := concatnormal("a", "b")
		_ = ans
	}
}
func BenchmarkFmtSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ans := concatfmt("a", "b")
		_ = ans
	}
}

func BenchmarkStringsBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ans := concatstrings("a", "b")
		_ = ans
	}
}

func BenchmarkStats(b *testing.B) {
	b.ReportAllocs() // u can skip this if u mention -benchmem in command while testing
	for i := 0; i < b.N; i++ {
		ans := benchStatTest("a", "b")
		_ = ans
	}
}
