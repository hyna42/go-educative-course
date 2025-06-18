package main

// import (
// 	"bufio"
// 	"fmt"
// 	"io"
// 	"os"
// )

func main() {
	//Reading & Writing using Scanln /  Sscan-family
	// var (
	// 	firstName, lastName string
	// )

	// fmt.Println("Please enter your full name: ")
	// fmt.Scanln(&firstName, &lastName)
	// fmt.Printf("Hi you %s %s!\n", firstName, lastName)

	//Reading & Writing using bufio
	// var inputReader *bufio.Reader
	// var (
	// 	input2 string
	// 	err    error
	// )
	// inputReader = bufio.NewReader(os.Stdin)
	// fmt.Println("Please enter some input: ")

	// input2, err = inputReader.ReadString('\n')

	// if err == nil {
	// 	fmt.Printf("Hi you again : %s\n", input2)
	// }
	//challenge 1 : read and write
	// inputReader := bufio.NewReader(os.Stdin)
	// fmt.Println("Please enter some input, type S in the new line to stop: ")
	// for {
	// 	input, err := inputReader.ReadString('\n')
	// 	if err != nil {
	// 		fmt.Printf("An error occurred: %s\n", err)
	// 		return
	// 	}
	// 	if strings.TrimSpace(input) == "S" {
	// 		fmt.Println("Here are the counts:")
	// 		fmt.Printf("Number of characters: %d\n", nrchars)
	// 		fmt.Printf("Number of words: %d\n", nrwords)
	// 		fmt.Printf("Number of lines: %d\n", nrlines)
	// 		os.Exit(0)
	// 	}
	// 	Counters(input)
	// }

	//reading file
	
	//reading with readfile
	var file = "read.txt"
	// RwithBufferIo(file)
	// RwithReadFile(file)

}
