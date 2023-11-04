package service

import (
	"testing"

	"github.com/shiva0612/go-learnings/testing-framework/using_mockery/mocks"
)

func TestNotify(t *testing.T) {
	service := new(mocks.Msgservice)
	service.On("Sendmsg", "order placed").Return(true)

	e := ecom{msgservice: service}
	e.Notify("order placed")
	service.AssertExpectations(t)
}
