package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func bufioReader() {

	b := bufio.NewReaderSize(strings.NewReader("shiva-surya-shiva-surya"), 17)
	fmt.Println(b.Buffered(), b.Size())

	by := make([]byte, 17)
	b.Read(by)
	fmt.Println(string(by))

	b1, _ := b.ReadByte()
	fmt.Println(string(b1))
	fmt.Println(b.Buffered(), b.Size())

	b2, _ := b.ReadByte()
	fmt.Println(string(b2))
	b.WriteTo(os.Stdout)
}
func bufr(f *os.File) {
	dest, _ := os.Open("destination file")
	bfr := bufio.NewReaderSize(f, 10) //intialize reader whic read atleast 10 bytes on each Read() call

	//it will read 3 bytes
	b := make([]byte, 3)
	for {
		n, err := bfr.Read(b)
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		fmt.Println(b[:n])
	}

	b = make([]byte, 20)
	for {
		// n, err := bfr.Read(b) -> this way it will only read 10 bytes, but u wanted to 20
		n, err := io.ReadFull(bfr, b) //use this
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		fmt.Println(b[:n])
	}

	//does not buffer since f(os.file) has readfrom method so (f -> dest) nobuffering
	bfr.WriteTo(dest)
}

func bufw(f *os.File) {
	src, _ := os.Open("source.file")
	bufw := bufio.NewWriterSize(f, 10) //intialize writer which buffers exactly 10 bytes and writes to underlying writer only if buffer is exhausted

	// since f(os.file) has readfrom method so (f <- src)
	bufw.ReadFrom(src)

	/*
		this is does not write anything to stdout -> it will be buffered and on exit it will be lost
			bufio.newWriterSize(os.stdout, 10)
			bufw.write(2)
			exit

		content is written to stdout, since flush is called
			bufio.newWriterSize(os.stdout, 10)
			bufw.write(2)
			bufw.flush()
			exit

		directly written to stdout without buffering since 12 > 10
			bufio.newWriterSize(os.stdout, 10)
			bufw.write(12)
			exit

			bufio.newWriterSize(os.stdout, 10)
			bufw.write(10) //will buffer
			bufw.write(2) //prev 10 bytes will be printed and 2 bytes will be buffered
			exit

	*/

}
