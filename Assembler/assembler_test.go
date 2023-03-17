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

func TestConversionErrors(t *testing.T) {
	inst2Long := []string{"push 1 2"}
	OTypeLong := []string{"add 1"}
	ItypeShort := []string{"addi"}
	invalidInst := []string{"elvis"}
	imm2Big := []string{"push 1111"}

	c := NewConversion(inst2Long)
	assert.Error(t, c.ToBinary("test"), "too many arguments")

	c = NewConversion(OTypeLong)
	assert.Error(t, c.ToBinary("test"), "add does not take an immediate value")

	c = NewConversion(ItypeShort)
	assert.Error(t, c.ToBinary("test"), "addi requires an immediate value")

	c = NewConversion(invalidInst)
	assert.Error(t, c.ToBinary("test"), "elvis is an invalid instruction")

	c = NewConversion(imm2Big)
	assert.Error(t, c.ToBinary("test"), "1111 is too large for an immediate value")

}

func TestAssemblerSimple(t *testing.T) {
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
			incorrectLines++
			t.Errorf("Expected %s, got %s", correctLines[i], line)
		}
	}

	// write the output to a file
	f, err := os.Create("test_result.out")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	for _, line := range result {
		f.WriteString(line + "\n")
	}

	assert.Equal(t, 0, incorrectLines)
}
