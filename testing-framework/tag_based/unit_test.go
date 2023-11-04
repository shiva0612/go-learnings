//go:build unit
// +build unit

package main

import "testing"

func TestOne(t *testing.T) {
	logger("one")
}

func TestTwo(t *testing.T) {
	logger("two")
}

func TestThree(t *testing.T) {
	logger("three")
}
