package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func ioTest() {
	// bytesBuffer()
	// stringsReader()
	bytesReader()
}
func bytesReader() {
	br := bytes.NewReader([]byte("shiva-surya"))
	br.WriteTo(os.Stdout)
	fmt.Println()
	fmt.Println(br.Size(), br.Len())
	for i := 0; i < 5; i++ {
		br.UnreadByte()
	}
	fmt.Println(br.Size(), br.Len())
	br.WriteTo(os.Stdout)
	fmt.Println(br.Size(), br.Len())
}

func stringsReader() {
	sb := strings.NewReader("shiva-surya")
	sb.WriteTo(os.Stdout)
	fmt.Println()
	fmt.Println(sb.Size(), sb.Len())
	for i := 0; i < 5; i++ {
		sb.UnreadByte()
	}
	fmt.Println(sb.Size(), sb.Len())
	sb.WriteTo(os.Stdout)
	fmt.Println(sb.Size(), sb.Len())
}
func bytesBuffer() {
	bb := bytes.NewBufferString("shiva")
	bb.ReadFrom(bytes.NewBufferString("some other reader"))
	fmt.Println(bb.Cap(), bb.Len())
	// bb.Reset()
	bb.WriteTo(os.Stdout)
	fmt.Println()
	fmt.Println(bb.Cap(), bb.Len())
	for i := 0; i < 10; i++ {
		bb.UnreadByte()
	}
	fmt.Println(bb.Cap(), bb.Len())
	bb.WriteTo(os.Stdout)
	fmt.Println(bb.Cap(), bb.Len())
}

/*
type Builder
    func (b *Builder) Cap() int #cap doubles when we keep on writing
    func (b *Builder) Grow(n int)
    func (b *Builder) Len() int #actual length of the underlying string

    func (b *Builder) Reset() #both len and cap is 0
    func (b *Builder) String() string

    func (b *Builder) Write(p []byte) (int, error)
    func (b *Builder) WriteByte(c byte) error
    func (b *Builder) WriteRune(r rune) (int, error)
    func (b *Builder) WriteString(s string) (int, error)
------------------------------------------------------------------------------------------------------------------
type Reader (strings/bytes reader is same)
    func NewReader(s string) *Reader

    func (r *Reader) Len() int #len of unread bytes
    func (r *Reader) Size() int64 #size of the underlying string

    func (r *Reader) Reset(s string) #the size and the cap will be reset according to new string passed
    func (r *Reader) WriteTo(w io.Writer) (n int64, err error) #will write to writer, but still u can unread and start reading again


    func (r *Reader) Read(b []byte) (n int, err error)
    func (r *Reader) ReadAt(b []byte, off int64) (n int, err error)
    func (r *Reader) ReadByte() (byte, error)
    func (r *Reader) ReadRune() (ch rune, size int, err error)
    func (r *Reader) Seek(offset int64, whence int) (int64, error)

    func (r *Reader) UnreadByte() error
    func (r *Reader) UnreadRune() error
------------------------------------------------------------------------------------------------------------------
bytes buffer

    this has a internal []byte
    all write methods keep increasing the size of the above
    also, readfrom(some reader) -> just increases the size of the byte

    when u start reading bytes.buffer
        this will just keep moving the pointer and will u the bytes

    when reset, writeTo(os.Stdout) -> the content will be gone , but the size of []byte will remain same
        but, in strings reader & bytes reader -> while writeTo() -> only the pointer is moved, we can unreadByte(), and read again

func io() {
	bb := bytes.NewBufferString("shiva")
	bb.ReadFrom(bytes.NewBufferString("some other reader"))
	fmt.Println(bb.Cap(), bb.Len())
	// bb.Reset()
	bb.WriteTo(os.Stdout)
	fmt.Println(bb.Cap(), bb.Len())
}
*/
