package Emulator

import (
	"Assembler"
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimplePrograms(t *testing.T) {
	proc := NewProcessor()
	p := Assembler.NewParser("/test_files/add.txt")
	pErr := p.Parse()
	if pErr != nil {
		t.Errorf("Error parsing file: %s", pErr)
	}

	c := Assembler.NewConversion(p.GetLines())
	cErr := c.ToBinary("")
	if cErr != nil {
		t.Errorf("Error converting file: %s", cErr)
	}

	program := make([]int16, len(c.GetOutput()))
	for i, line := range c.GetOutput() {
		num, err := strconv.ParseInt(line, 2, 16)
		fmt.Println(line, num)
		if err != nil {
			fmt.Println("Error converting binary string to int16")
			fmt.Println(err)
		}
		program[i] = int16(num)
	}
	proc.LoadProgram(program)

	for !proc.ProgramExit {
		proc.DecodeInstruction()
	}

	// fmt.Println(proc.ram[proc.wsp])
	assert.Equal(t, int16(30), proc.ram[proc.wsp])
}
