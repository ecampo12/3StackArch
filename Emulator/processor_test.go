package Emulator

import (
	"Assembler"
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Helper function that assembles a program from a file and returns the binary
func assemblerProgram(fileName string) ([]uint16, error) {
	p := Assembler.NewParser(fileName)
	pErr := p.Parse()
	if pErr != nil {
		fmt.Errorf("Error parsing file: %s", pErr)
	}

	c := Assembler.NewConversion(p.GetLines())
	cErr := c.ToBinary("")
	if cErr != nil {
		fmt.Errorf("Error converting file: %s", cErr)
	}

	program := make([]uint16, len(c.GetOutput()))
	for i, line := range c.GetOutput() {
		num, err := strconv.ParseUint(line, 2, 16)
		// fmt.Println(line, num)
		if err != nil {
			fmt.Println("Error converting binary string to int16")
			fmt.Println(err)
		}
		program[i] = uint16(num)
	}
	return program, nil
}

// Just a sanity check to make sure I didn't miss any opcodes
func TestDecodeInstructions(t *testing.T) {
	proc := NewProcessor()
	ItypeInstructions, err := assemblerProgram("/test_files/Itype.txt")

	if err != nil {
		fmt.Println(err)
	}
	proc.LoadProgram(ItypeInstructions)
	proc.Run()

	proc = NewProcessor()
	OtypeInstructions, err := assemblerProgram("/test_files/Otype.txt")

	if err != nil {
		fmt.Println(err)
	}
	proc.LoadProgram(OtypeInstructions)
	// pre-emptively set the stack pointers to prevent errors
	proc.wsp = 0x02D0    // 720
	proc.ram[0x024F] = 8 // 591
	proc.ram[726] = 4
	proc.Run()
}

func TestSimplePrograms(t *testing.T) {
	proc := NewProcessor()

	addProgram, err := assemblerProgram("/test_files/add.txt")
	if err != nil {
		fmt.Println(err)
	}
	proc.LoadProgram(addProgram)

	proc.Run()

	// fmt.Println(proc.ram[proc.wsp])
	assert.Equal(t, uint16(30), proc.ram[proc.wsp])

	proc = NewProcessor()

	subProgram, err := assemblerProgram("/test_files/sub.txt")
	if err != nil {
		fmt.Println(err)
	}
	proc.LoadProgram(subProgram)

	proc.Run()

	// fmt.Println(proc.ram[proc.wsp])
	assert.Equal(t, uint16(10), proc.ram[proc.wsp])
}
