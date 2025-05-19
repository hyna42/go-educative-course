package main

func SumArr(arr *[4]int) (sum int) {
	for _, v := range arr {
		sum += v
	}
	return
}
