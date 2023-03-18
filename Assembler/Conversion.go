package assembler

import (
	"errors"
	"fmt"
	"math/bits"
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

func (c *Conversion) ToBinary(sType string) error {
	for _, line := range c.lines {
		convert, err := c.convert(line, sType)
		if err != nil {
			// fmt.Errorf("Error converting %s: %s", line, err)
			return err
		}
		c.output = append(c.output, convert)
	}
	return nil
}

func (c *Conversion) convert(line string, sType string) (string, error) {
	var opcode string
	var imm string
	split := strings.Split(line, " ")

	if len(split) > 2 {
		return "", errors.New("too many arguments")
	}

	opcode = Instructions().NameToBinary[split[0]]
	if opcode == "" {
		return "", errors.New(split[0] + " is an invalid instruction")
	}

	switch Instructions().instructionType[split[0]] {
	case 0:
		if len(split) > 1 {
			return "", errors.New(split[0] + " does not take an immediate value")
		}
		imm = "0000000000"
	case 1:
		if len(split) < 2 {
			return "", errors.New(split[0] + " requires an immediate value")
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
			return "", errors.New(split[0] + " immediate value too large. Max value is 1023")
		}

		numBits := bits.UintSize
		binaryRep := bits.RotateLeft16(uint16(temp), numBits)

		binaryStr := fmt.Sprintf("%010b", binaryRep)
		if len(binaryStr) > 10 {
			imm = binaryStr[len(binaryStr)-10:]
		} else {
			// Pad with leading zeros if necessary
			imm = fmt.Sprintf("%010s", binaryStr)
		}
		// fmt.Println(opcode + " " + imm)
	}

	if sType == "test" {
		return opcode + " " + imm, nil
	}

	return opcode + imm, nil

	// hex := opcode + imm
	// // string to int
	// // num, _ := strconv.Atoi(hex)
	// num, _ := strconv.ParseInt(hex, 2, 16)
	// fmt.Println(num)
	// fmt.Printf("%04X\n", num)
	// // convert to hex
	// // return strconv.FormatInt(int64(num), 16), nil

	// return fmt.Sprintf("%04x", num), nil
}
