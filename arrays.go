package main

func SumArr(arr *[4]int) (sum int) {
	for _, v := range arr {
		sum += v
	}
	return
}

//challenge : Filling Array with Loop Counter

func LoopArray() [15]int {
	var arr [15]int
	for i := range arr {
		arr[i] = i
	}
	return arr
}

// Challenge: Finding Fibonacci Numbers with Array

func Fibs() [10]int64 {
	var fib [10]int64

	for i := range len(fib) {
		if i <= 1 {
			fib[i] = 1
		} else {
			fib[i] = fib[i-1] + fib[i-2]
		}
	}
	return fib
}
