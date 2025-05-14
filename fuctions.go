package main

func Greeting(name *string) {
	println("In greeting: Hi first ", *name)
	*name = "Johnny"
	println("In greeting: Hi again", *name)
}

// challenge
func SumProductDiff(i, j int) (s int, p int, d int) { // named version

	// s = i + j
	// p = i * j
	// d = i - j

	// return s, p, d

	s, p, d = i+j, i*j, i-j
	return
}
