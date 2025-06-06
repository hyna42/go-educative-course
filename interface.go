package main

import "fmt"

// Challenge: Make a Simple Interface
type Simpler interface {
	Get() int
	Set(p int)
}

type Simple struct {
	v int
}


func (s *Simple) Get() int {
	return s.v
}

func (s *Simple) Set(u int) {
	s.v = u
}

func operate(s Simpler) {
	s.Set(42)
	fmt.Println("Value:", s.Get())
}
