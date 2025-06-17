package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//Reading & Writing using Scanln /  Sscan-family
	var (
		firstName, lastName string
		// i                      int
		// f                      float32
		// input1                 = "56.12 / 5212 / Go"
		// format                 = "%f / %d / %s"
	)

	fmt.Println("Please enter your full name: ")
	fmt.Scanln(&firstName, &lastName)
	fmt.Printf("Hi you %s %s!\n", firstName, lastName)
	// fmt.Sscanf(input1, format, &f, &i, &s)
	// fmt.Println("From the string we read: ", f, i, s)

	//Reading & Writing using bufio
	var inputReader *bufio.Reader
	var (
		input2 string
		err    error
	)
	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("Please enter some input: ")

	input2, err = inputReader.ReadString('\n')

	if err == nil {
		fmt.Printf("Hi you again : %s\n", input2)
	}

}
