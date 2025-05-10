// Challenge 2 : Control structures
package main

import "fmt"

//switch
func Season(month int) string {
	var season string
	switch month {
	case 1, 2, 12:
		season = "Winter"
	case 3, 4, 5:
		season = "Spring"
	case 6, 7, 8:
		season = "Summer"
	case 9, 10, 11:
		season = "Autumn"
	default:
		season = "Season unknown"
	}
	return season
}

//break and continue
func FizzBuzz() {
	for i := 0; i <= 60; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Print("FizzBuzz")
			continue
		}
		if i%3 == 0 {
			fmt.Print("Fizz")
			continue
		}
		if i%5 == 0 {
			fmt.Print("Buzz")
			continue
		}

		fmt.Print(i)
	}
}
