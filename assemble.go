package main

import (
	"Assembler"
	"fmt"
	"os"
)

// explains how the assembler is used
func Usage() {
	fmt.Println("Usage: assembler <inputFile> <outputFile>")
}

func main() {
	// check for correct number of arguments
	if len(os.Args) != 3 {
		Usage()
		os.Exit(1)
	}

	// get the input and output file names
	inputFile := os.Args[1]
	outputFile := os.Args[2]

	err := Assembler.Assemble(inputFile, outputFile)
	if err != nil {
		fmt.Println("Error assembling:", err)
		os.Exit(1)
	}
}
