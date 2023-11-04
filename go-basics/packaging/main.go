package main

import (
	"github.com/shiva0612/go-learnings/go-basics/packaging/database"
	db "github.com/shiva0612/go-learnings/go-basics/packaging/database/db"
	db2 "github.com/shiva0612/go-learnings/go-basics/packaging/database/db2"
)

/*
	if the pkg is of v1
	then simple

	go get pkgName
	import ("pkgName")
	---

	if the pkg is v2 or higher

	go get pkgName/v2
	import("pkgName/v2")
*/

var (
	//import module & main.Main_var
	Main_var = "main var"
)

func main() {
	_ = db.Name1
	_ = database.Name
	_ = db2.Name2
}
