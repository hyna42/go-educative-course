package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var nrchars, nrwords, nrlines = 0, 0, 0

func Counters(input string) {
	nrchars += len(input) - 2
	nrwords += strings.Count(input, " ") + 1
	nrlines++
}

// methode 1 :with buffor.io
func RwithBufferIo(file string) {
	inputFile, err := os.Open(file)
	if err != nil {
		fmt.Println("Error on reading file", err)
		return
	}
	defer inputFile.Close()
	reader := bufio.NewReader(inputFile)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			fmt.Println("R BufferIO ==>",line)
			break
		}
		fmt.Println("Erreur :", err)
	}

}

// methode 2 readfile
func RwithReadFile(file string) {
	content, er := os.ReadFile(file)
	if er != nil {
		fmt.Println("Error on reading file")
	}
	fmt.Println("R ReadFile ==>", string(content))
}
