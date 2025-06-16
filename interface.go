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
type Square struct {
	side float32
}

type Triangle struct { // implement this struct
	base   float32
	height float32
}

type AreaInterface interface {
	Area() float32
}

type PeriInterface interface { // implement this interface
	Perimeter() float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func (sq *Square) Perimeter() float32 { // implement method called on square to calculate its perimeter
	return sq.side * 4
}

func (tr *Triangle) Area() float32 { // implement method called on triangle to calculate its area
	return 0.5 * (tr.base * tr.height)
}
