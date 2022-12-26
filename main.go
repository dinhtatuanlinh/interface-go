package main

import (
	"fmt"
)

func main() {
	a := newInstance()
	b := foo{
		model: a,
	}
	b.model.run()
	// method 2
	var naa boo
	naa = &noo{}
	naa.run()
}

func newInstance() *noo {
	return &noo{}
}

type boo interface {
	run()
}

type noo struct {
	boo
}

func (n *noo) run() {
	fmt.Println("running")
}

type foo struct {
	model boo
}
