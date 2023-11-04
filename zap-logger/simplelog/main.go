package main

import (
	"fmt"
	"log"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

type DB_ERR struct {
	code  int
	level string
	msg   string
}

func (de DB_ERR) Error() string {
	return fmt.Sprintf("level:%s	code:%d	msg:%s", de.level, de.code, de.msg)
}
func getDB_ERR(code int, level string, msg string) DB_ERR {
	return DB_ERR{code, level, msg}
}

func main() {
	formatted_error := fmt.Errorf("this is simple formatted error")
	log.Println(formatted_error.Error())

	err := getDB_ERR(1, "tolerable", "this is error msg")
	log.Println(err.Error())
}
