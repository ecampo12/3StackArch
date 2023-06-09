package Assembler

type Instruction struct {
	NameToBinary    map[string]string
	instructionType map[string]int
}

func Instructions() *Instruction {
	return &Instruction{
		NameToBinary: map[string]string{
			"add":    "100000",
			"addi":   "100001",
			"and":    "100010",
			"andi":   "100100",
			"blt":    "101000",
			"bgt":    "110000",
			"beq":    "110001",
			"call":   "110010",
			"cmp":    "111010",
			"clr":    "110111",
			"j":      "101010",
			"js":     "101110",
			"ld":     "101111",
			"lui":    "100101",
			"mdsp":   "100110",
			"ori":    "000001",
			"or":     "000010",
			"pop":    "000100",
			"pushi":  "000110",
			"push":   "001000",
			"ret":    "001010",
			"sfl":    "001110",
			"sfr":    "010000",
			"st":     "010010",
			"sub":    "010011",
			"swap":   "011000",
			"exit":   "000000",
			"ldinbf": "111100",
			"iret":   "111101",
			"iexc":   "111000",
			"eret":   "111001",
			"ignr":   "111010",
		},
		instructionType: map[string]int{
			// 0 = O, 1 = I
			"add":    0,
			"addi":   1,
			"and":    0,
			"andi":   1,
			"blt":    1,
			"bgt":    1,
			"beq":    1,
			"call":   1,
			"cmp":    0,
			"clr":    1,
			"j":      1,
			"js":     0,
			"ld":     1,
			"lui":    1,
			"mdsp":   1,
			"ori":    1,
			"or":     0,
			"pop":    1,
			"pushi":  1,
			"push":   1,
			"ret":    0,
			"sfl":    1,
			"sfr":    1,
			"st":     1,
			"sub":    0,
			"swap":   0,
			"exit":   0,
			"ldinbf": 0,
			"iret":   0,
			"iexc":   0,
			"eret":   0,
			"ignr":   0,
		},
	}
}
