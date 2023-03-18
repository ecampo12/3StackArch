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

// helper function to write test results to a file, creates a directory if it doesn't exist
func writeFile(fileName string, lines []string) {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// create the directory if it doesn't exist
	if _, err := os.Stat(wd + "/test_results"); os.IsNotExist(err) {
		os.Mkdir(wd+"/test_results", 0755)
	}

	// create the file in the directory
	file, err := os.Create(wd + "/test_results/" + fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		_, err := w.WriteString(line + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
	w.Flush()
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

func TestParserError(t *testing.T) {
	p := NewParser("/test_files/long_program.txt")
	assert.Error(t, p.Parse(), "file exceeds 256 lines")
}

func TestAssemblerSimple(t *testing.T) {
	incorrectLines := 0
	correctLines := readFile("/test_files/Simple_Instructions_trans.txt")

	parser := NewParser("/test_files/Simple_Instructions.txt")

	parser.Parse()
	writeFile("simple_test_result.parse", parser.GetLines())

	c := NewConversion(parser.GetLines())

	err := c.ToBinary("test")
	assert.NoError(t, err, "should not have an error")
	result := c.GetOutput()

	assert.NotEmpty(t, result, "result should not be empty")
	for i, line := range result {
		if line != correctLines[i] {
			incorrectLines++
			t.Errorf("Expected %s, got %s", correctLines[i], line)
		}
	}

	writeFile("simple_test_result.out", result)
	assert.Equal(t, 0, incorrectLines)
}

func TestAssemblerSimpleLables(t *testing.T) {
	incorrectLines := 0
	correctLines := readFile("/test_files/Simple_w_Lables_trans.txt")

	parser := NewParser("/test_files/Simple_w_Lables.txt")
	parser.Parse()
	writeFile("simple_w_lables_test_result.parse", parser.GetLines())

	c := NewConversion(parser.GetLines())
	err := c.ToBinary("test")
	assert.NoError(t, err, "should not have an error")
	result := c.GetOutput()

	assert.NotEmpty(t, result, "result should not be empty")

	for i, line := range result {
		if line != correctLines[i] {
			incorrectLines++
			t.Errorf("Expected %s, got %s", correctLines[i], line)
		}
	}

	writeFile("simple_w_lables_test_result.out", result)
	assert.Equal(t, 0, incorrectLines)
}

func TestAssemblerComplex(t *testing.T) {
	incorrectLines := 0
	correctLines := readFile("/test_files/Real_Prime_trans.txt")

	parser := NewParser("/test_files/Real_Prime.txt")
	parser.Parse()
	writeFile("real_prime_test_result.parse", parser.GetLines())

	c := NewConversion(parser.GetLines())
	err := c.ToBinary("test")
	assert.NoError(t, err, "should not have an error")
	result := c.GetOutput()

	assert.NotEmpty(t, result, "result should not be empty")

	for i, line := range result {
		if line != correctLines[i] {
			incorrectLines++
			t.Errorf("Expected %s, got %s", correctLines[i], line)
		}
	}

	writeFile("real_prime_test_result.out", result)
	assert.Equal(t, 0, incorrectLines)
}
