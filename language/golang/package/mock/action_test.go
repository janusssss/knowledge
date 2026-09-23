package mock

import "testing"

var nameP func() string = name

func TestCall(t *testing.T) {
	nameP = func() string {
		return "mock"
	}
	call()
}
