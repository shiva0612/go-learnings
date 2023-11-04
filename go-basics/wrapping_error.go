package main

import (
	"errors"
	"fmt"
)

var myerror = errors.New("myerror")

func wrapping_error() {
	err := geterr(1)
	if errors.Is(err, myerror) {
		fmt.Println(err.Error())                //myerror: given 1
		fmt.Println(errors.Unwrap(err).Error()) //myerror
		return
	}
}

func geterr(i int) error {

	if i == 1 {
		return fmt.Errorf("%w: given 1", myerror)
	}
	return fmt.Errorf("normal error")
}
