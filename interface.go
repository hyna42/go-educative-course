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

func Operate(s Simpler) {
	s.Set(42)
	fmt.Println("Value:", s.Get())
}

//Challenge: Advancing the Simple Interface
type RSimple struct {
	r int
}

func (r *RSimple) Get() int {
	return r.r
}

func (r *RSimple) Set(v int) {
	r.r = v
}

func ROperate(s Simpler) {
	s.Set(42)
	fmt.Println("Value:", s.Get())

	switch v := s.(type) {
	case *Simple:
		fmt.Println("Handled a Simple struct:", v.v)
	case *RSimple:
		fmt.Println("Handled a RSimple struct:", v.r)
	default:
		fmt.Println("Unknown Simpler type")
	}
}

// Challenge: Advancing the Shapes Analysis