package main

import "fmt"

func main() {
	name := "amadou"
	Greeting(&name)

	//challenge functon : multiple return
	var s,p,d int = SumProductDiff(3, 4)

	fmt.Printf("Sum : %d, Prod : %d, Diff : %d ",s,p,d)

}
