package service

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"
)

type smsmock struct {
	mock.Mock
}

func (sm *smsmock) sendmsg(msg string) bool {
	fmt.Println("faking sending msg")
	return_args := sm.Called(msg)

	fmt.Println("args = ", return_args) //args=true in TestNotify, args=1, "nothing" in testnotify2
	return return_args.Bool(0)
}

func TestNotify(t *testing.T) {
	service := new(smsmock)
	service.On("sendmsg", "order placed").Return(true)

	e := ecom{msgservice: service}
	e.Notify("order placed")
	service.AssertExpectations(t)
}
func TestNotify2(t *testing.T) {
	service := new(smsmock)
	service.On("sendmsg", "order placed").Return(true, 1, "nothing")

	e := ecom{msgservice: service}
	e.Notify("order placed")
	service.AssertExpectations(t)
}
