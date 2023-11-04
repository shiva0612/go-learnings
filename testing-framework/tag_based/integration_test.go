//go:build integration
// +build integration

package main

import "testing"

func TestFour(t *testing.T) {
	logger("Four")
}

func TestFive(t *testing.T) {
	logger("Five")
}

func TestSix(t *testing.T) {
	logger("Six")
}
