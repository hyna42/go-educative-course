package main

import "strings"

var nrchars, nrwords, nrlines = 0, 0, 0

func Counters(input string) {
	nrchars += len(input) - 2 
	nrwords += strings.Count(input, " ") + 1
	nrlines++
}
