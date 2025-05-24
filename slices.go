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

//Challenge: Filtering with Higher-Order Functions
func FilterSlices(s []int, fn func(int) bool) []int {
	var p []int

	for _, v := range s {
		if fn(v) {
			p = append(p, v)
		}
	}
	return p
}

//Challenge: Bubble Sort the Slice

func BubbleSort(s []int) []int {

	n := len(s)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
	return s
}

//Challenge: Reverse a String
func ReverseString(s string) string {
	sl2 := []byte(s)
	for i, j := 0, len(sl2)-1; i < j; i, j = i+1, j-1 {
		sl2[i], sl2[j] = sl2[j], sl2[i]
	}
	return string(sl2)
}

