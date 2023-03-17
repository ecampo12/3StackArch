package assembler

import (
	"bufio"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// helper function to read a file and return the contents as a string slice
func readFile(fileName string) []string {
	var lines []string
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	file, err := os.Open(wd + fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

func TestAssembler(t *testing.T) {
	// var testLines []string
	incorrectLines := 0
	correctLines := readFile("/test_files/Simple_Instructions_trans.txt")

	parser := NewParser("/test_files/Simple_Instructions.txt")

	parser.Parse()

	c := NewConversion(parser.GetLines())

	c.ToBinary("test")
	result := c.GetOutput()

	for i, line := range result {
		if line != correctLines[i] {
			t.Errorf("Expected %s, got %s", correctLines[i], line)
		}
	}

	assert.Equal(t, incorrectLines, 0)
}
