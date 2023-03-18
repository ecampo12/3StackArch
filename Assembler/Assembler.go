package Assembler

import (
	"bufio"
	"fmt"
	"os"
)

// explains how the assembler is used
func Usage() {
	fmt.Println("Usage: assembler </inputFile> <outputFile>")
}

func Assemble(inputFile string, outputFile string) error {
	// open the input file
	wd, err := os.Getwd()
	if err != nil {
		// fmt.Println("Error getting working directory:", err)
		// os.Exit(1)
		return err
	}
	file, err := os.Open(wd + inputFile)
	// file, err := os.Open(inputFile)
	fmt.Println(wd)
	fmt.Println(wd + inputFile)
	if err != nil {
		// fmt.Println("Error opening input file:", err)
		// os.Exit(1)
		return err
	}
	defer file.Close()

	// parse the input file
	parser := NewParser(inputFile)
	parser.Parse()

	// convert the parsed input to binary
	c := NewConversion(parser.GetLines())
	cErr := c.ToBinary("")
	if cErr != nil {
		// fmt.Println("Error converting to binary:", err)
		// os.Exit(1)
		return cErr
	}
	result := c.GetOutput()

	// write the binary output to the output file
	file, err = os.Create(wd + "/" + outputFile)
	if err != nil {
		// fmt.Println("Error creating output file:", err)
		// os.Exit(1)
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range result {
		_, err := w.WriteString(line + "\n")
		if err != nil {
			// fmt.Println("Error writing to output file:", err)
			// os.Exit(1)
			return err
		}
	}

	w.Flush()

	return nil
}
