package main

import (
	"os"
)

func main() {

	// tee()
	// pipe()
	multiread()
	multiwrite()
	f, _ := os.Open("file.mov")
	bufr(f)
	bufioReader()
	bufw(f)
}
