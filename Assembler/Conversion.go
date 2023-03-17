package assembler

import (
	"fmt"
	"strconv"
	"strings"
)

type Conversion struct {
	lines  []string
	output []string
}

func NewConversion(lines []string) *Conversion {
	return &Conversion{
		lines: lines,
	}
}

func (c *Conversion) GetOutput() []string {
	return c.output
}

func (c *Conversion) ToBinary(sType string) {
	for _, line := range c.lines {
		c.output = append(c.output, c.convert(line, sType))
	}
}

func (c *Conversion) convert(line string, sType string) string {
	var opcode string
	var imm string
	split := strings.Split(line, " ")

	if len(split) > 2 {
		// TODO: Handle this error
	}

	// TODO: add check to see if the instruction is valid
	opcode = Instructions().NameToBinary[split[0]]

	switch Instructions().instructionType[split[0]] {
	case 0:
		if len(split) > 1 {
			// TODO: throrw error if there is more than 1 argument
		}
		imm = "0000000000"
	case 1:
		if len(split) < 2 {
			// TODO: throw error if there are less than 2 arguments
		}
		var temp int64

		// check the base of the immediate value
		if len(split[1]) > 2 {
			if split[1][0:2] == "0x" {
				temp, _ = strconv.ParseInt(split[1][2:], 16, 32)
			} else if split[1][0:2] == "0b" {
				temp, _ = strconv.ParseInt(split[1][2:], 2, 32)
			} else {
				temp, _ = strconv.ParseInt(split[1], 10, 32)
			}
		} else {
			temp, _ = strconv.ParseInt(split[1], 10, 32)
		}

		if temp > 1023 {
			// TODO: throw error if the immediate value is too large
		}

		if temp >= 0 {
			// imm = strconv.FormatInt(temp, 2)
			imm = fmt.Sprintf("%0"+strconv.Itoa(int(temp))+"d", temp)
		} else {
			imm = strconv.Itoa(int(temp))
			imm = imm[len(imm)-10:]
		}
	}

	if sType == "test" {
		return opcode + " " + imm
	}
	if sType == "bin" {
		return opcode + imm
	}

	test := opcode + imm
	// string to int
	num, _ := strconv.Atoi(test)
	// convert to hex
	return strconv.FormatInt(int64(num), 16)
}
