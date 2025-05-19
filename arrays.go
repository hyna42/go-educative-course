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
