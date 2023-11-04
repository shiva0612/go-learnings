```
byte = uint8
rune = int32

string is a list of bytes but not rune
but while looping string we get rune 
all the string methods gives answer in the context of bytes

msg := "a🐜"
n1 := len(msg) // 1+4=5 // every string in computer is stored in bytes
n2 := utf8.RuneCountInString(msg) // 2

---looping over string gives rune 
for _,i := range "a🐜"{
i -> type rune -> some int32 number -> 97,128028

}

func main() {
	a := "a🐧a🐧"
	for i, v := range a {
		fmt.Println(i, v)
	}
}
//output -> indicated that looping is by rune, see the index i=0156
0 97
1 128039
5 97
6 128039

------string to bytes/rune

name := "shiva"
b := []byte(name)
r := []rune(name)

------byte/rune to string

string([]byte)
string([]rune)
```
