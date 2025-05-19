package main

import "fmt"

func main() {
	name := "amadou"
	Greeting(&name)

	//challenge functon : multiple return
	var s, p, d int = SumProductDiff(3, 4)

	fmt.Printf("Sum : %d, Prod : %d, Diff : %d\n ", s, p, d)

	sum := SumInts(1, 2, 3, 4, 888)
	//challenge sumInts
	fmt.Printf("Sum 2: %d\n", sum)

	//challenge factoriel
	fac := MyFactorial(10)
	fmt.Printf("Factorielle : %d\n", fac)

	//challenge Filter Even and Odd Numbers
	sl := []int{1, 2, 3, 4, 5, 7}

	even, odd := Filter(sl, IsEven)

	fmt.Printf("even : %d\n", even)
	fmt.Printf("odd : %d\n", odd)

	//arrays
	var arr [4]int

	for i := range len(arr) {
		arr[i] = i + 1
		fmt.Printf("Item at index %d is %d\n", i, arr[i])
	}

	//pass an array as parameter of function
	sumArr := SumArr(&arr)
	fmt.Printf("SumArr: %d\n", sumArr)

	//challenge : Filling Array with Loop Counter
	arr2 := LoopArray()
	fmt.Printf("Challenge Array 2 ==> : %d\n ", arr2)

}
