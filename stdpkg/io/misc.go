package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

/*
checkout io.copy, io.copyBuffer()
*/

// what ever u write x u can read on y (but not other way around)
func pipe() {

	pr, pw := io.Pipe()
	go func() {
		pw.Write([]byte("from pipe"))
		pw.Close()
	}()
	io.Copy(os.Stdout, pr)

}

// what ever u read from x write to y
func tee() {
	var r io.Reader = strings.NewReader("from tee")

	r = io.TeeReader(r, os.Stdout)

	// Everything read from r will be copied to stdout.
	if _, err := io.ReadAll(r); err != nil {
		log.Fatal(err)
	}
}

func multiread() {

	f, _ := os.Open("")

	header := strings.NewReader("<msg>")
	body_file, _ := os.Open("api_response.file")
	footer := strings.NewReader("</msg>")

	mr := io.MultiReader(header, body_file, footer)
	io.Copy(f, mr)

}
func multiwrite() {
	f1, _ := os.Create("file1")
	f2, _ := os.Create("file2")
	f3, _ := os.Create("file3")

	mw := io.MultiWriter(f1, f2, f3)

	//writes the below content into all writers
	fmt.Fprintln(mw, "line1")
	fmt.Fprintln(mw, "line2")

}
