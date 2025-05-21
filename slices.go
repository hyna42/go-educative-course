package main

//challenge : Finding Fibonacci Numbers with Slices

func FibArray(term int) []int {
	fib := make([]int, term)

	fib[0], fib[1] = 0, 1

	for i := 2; i < term; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	return fib
}

//challenge : magnify a slice
func Enlarge(s []int, factor int) []int {
	newLen := len(s) * factor
	enlarged := make([]int, newLen)

	copy(enlarged, s)
	return enlarged
}

// Challenge: Inserting Slice in a Slice
func InsertSlice(slice, insertion []string, index int) []string {
	result := make([]string, len(slice)+len(insertion))
	at := copy(result, slice[:index])
	at += copy(result[at:], insertion)
	copy(result[at:], slice[index:])
	return result

}
