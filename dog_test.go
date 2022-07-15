package main

import "testing"

func TestFoo_show(t *testing.T) {
	_check := check
	defer func() {
		check = _check
	}()
	check = func() bool {
		return false
	}
	f := Foo{}
	f.show()
}
