package main

import "reflect"

// this way of checking is only if a is interface
func checkType(a interface{}) {
	switch a.(type) {
	case string:
	case int:
	}
}

// use reflect pkg for straight forward checking
func checkType2(a string) bool {
	return reflect.TypeOf(a) == reflect.TypeOf("")
}
