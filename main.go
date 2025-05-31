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

	// Challenge: Finding Fibonacci Numbers with Array

	var arrFib = Fibs()
	fmt.Printf("Challenge Array 3 ==> : %d\n ", arrFib)

	//slices
	slice1 := arrFib[5:cap(arrFib)]
	fmt.Printf("Slice 1 ==> : %d\n ", slice1)

	slice2 := new([]int)

	slice3 := make([]int, 5)

	fmt.Printf("Type arrFib : %T\n", arrFib)
	fmt.Printf("Type slice1 : %T\n", slice1)

	fmt.Printf("Type slice2 : %T\n", slice2)

	fmt.Printf("Type slice3 : %T\n", slice3)

	// Finding Fibonacci Numbers with Slices
	fibArray := FibArray(5)
	fmt.Printf("Challenge slice 1 : %d\n", fibArray)

	//challenge : magnify a slice
	sliceTest := []int{1, 2, 3}
	fmt.Println("The length of s before enlarging is:", len(sliceTest))
	fmt.Println(sliceTest)

	sliceTest = Enlarge(sliceTest, 5)
	fmt.Println("The length of s after enlarging is:", len(sliceTest))
	fmt.Println(sliceTest)

	//challenge : inserting slice in a slide
	s1 := []string{"M", "N", "O", "P", "Q", "R"}
	s2 := []string{"A", "B", "C"}
	res := InsertSlice(s1, s2, 0)
	fmt.Println(res)
	res = InsertSlice(s1, s2, 3)
	fmt.Println(res)

	//Challenge : Filtering with Higher-Order Functions
	s3 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	res2 := FilterSlices(s3, IsEven)
	fmt.Println("Challenge slice 3", res2)

	//Challenge: Bubble Sort the Slice
	s4 := []int{4, 5, 2, 1, 3}
	res4 := BubbleSort(s4)
	fmt.Println("Challenge slice 4", res4)

	//Challenge: Reverse a String
	s5 := "Google"
	res5 := ReverseString(s5)
	fmt.Println("Last Challenge slice 5", res5)

	ps := make([]byte, 5) 
	ps = ps[2:4]

	//Challenge: Map the Days
	day := FindDay(4)
	fmt.Println("Day with index 4 : ",day)
}
