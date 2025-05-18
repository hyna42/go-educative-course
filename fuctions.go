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

func SumInts(list ...int) (sum int) {

	for _, v := range list {
		sum += v
	}
	return
}

//challenge : factorielle

func MyFactorial(n uint64) (fac uint64) {
	if n == 0 || n == 1 {
		fac = 1
	} else {
		fac = n * MyFactorial(n-1)
	}
	return
}
