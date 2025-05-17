package main

import "fmt"

func main() {
	name := "amadou"
	Greeting(&name)

	//challenge functon : multiple return
	var s,p,d int = SumProductDiff(3, 4)

	fmt.Printf("Sum : %d, Prod : %d, Diff : %d\n ",s,p,d)

	sum := SumInts(1,2,3,4,888)
	//challenge sumInts
	fmt.Printf("Sum 2: %d",sum)

}
