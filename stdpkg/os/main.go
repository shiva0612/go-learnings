package main

import (
	"fmt"
	"os"

	"github.com/subosito/gotenv"
)

func main() {
	// create_append_file()

	executable, _ := os.Executable() //while starting the program - just to be in sync
	dir, _ := os.Getwd()
	host, _ := os.Hostname()
	fmt.Println("host: ", host) //name of the machine as per the router
	fmt.Println("executable: ", executable)
	fmt.Println("dir: ", dir)
	// os.Getuid()
	// os.Getpid()

	working_with_env()

}

func working_with_env() {
	gotenv.Load(".env")
	// os.Getenv()
	// os.Clearenv()
	// os.Environ()
	// os.Setenv("NAME", "gopher")
	// os.Setenv("BURROW", "/usr/gopher")
	// os.Unsetenv()
	// fmt.Println(os.ExpandEnv("$NAME lives in ${BURROW}."))

}

/*
IMP ==== os.O_WRONLY: This flag specifies that the file should be opened in write-only mode. It allows writing data to the file but doesn't permit reading from it.

	dont forget this, specify whether to open file in read/write/both -> else you will get bad file descriptor error

os.O_APPEND: This flag specifies that the file should be opened in append mode. When writing to the file, the data will be added at the end of the file rather than overwriting existing content.
os.O_CREATE: This flag specifies that the file should be created if it doesn't exist. If the file already exists, this flag has no effect. When combined with other flags, it ensures that a new file is created when opening.
os.O_TRUNC: this is truncate the file content if present and starts writing fresh
*/
func create_append_file() {
	_, err := os.OpenFile("op.go", os.O_TRUNC|os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
}
