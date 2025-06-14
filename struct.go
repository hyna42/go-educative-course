package main

import (
	"fmt"
	"math"
)

//Challenge: Anonymous Struct

type S struct {
	a float32
	int
	string
}

// Challenge: Coordinates of a Point
type Point struct {
	X, Y float64
}

func (p *Point) Abs() float64 {

	return math.Sqrt(float64(p.X*p.X + p.Y*p.Y))
}

func (p *Point) Scale(s float64) {
	p.X = p.X * 5
	p.Y = p.Y * 5

	return
}

// Challenge: Make a Rectangle
type Rectangle struct {
	lenght int
	width  int
}

func (r *Rectangle) Area() int {
	return r.lenght * r.width
}

func (r *Rectangle) Perimeter() int {
	return 2 * (r.lenght + r.width)

}

// Challenge: Decide Employee Salary

type Employee struct {
	salary float64
}

func (this *Employee) GiveRaise(pct float64) float64 {

	return this.salary * (1 + pct)
}

// Challenge: Implement Stack Data Structure
const LIMIT = 4

type Stack struct {
	ix   int
	data [LIMIT]int
}

func (s *Stack) Push(n int) {
	if s.ix < LIMIT {
		s.data[s.ix] = n
		s.ix++
	}
}

func (s *Stack) Pop() int {
	if s.ix == 0 {
		return -1
	}
	s.ix--
	v := s.data[s.ix]
	s.data[s.ix] = 0
	return v

}

func (s *Stack) String() string {
	out := ""
	for i, v := range s.data {
		out += fmt.Sprintf("[%d:%d]", i, v)
	}
	return out
}
