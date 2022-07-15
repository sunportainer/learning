package main

import "fmt"

type Foo struct {
	name string
}

func (f *Foo) show() {
	if check() {
		fmt.Println("good boy")
	} else {
		fmt.Println("naughty boy")
	}

}

var check = func() bool {
	return true
}
