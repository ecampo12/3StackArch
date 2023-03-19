package Emulator

type Processor struct {
	wsp            int16 // work stack pointer
	dsp            int16 // data stack pointer
	rsp            int16 // return stack pointer
	programCounter int16 // program counter

	ram [0x03FF]int16 // 1024 words of memory

	// 16 flags:
	// 		15: CF - carry flag
	// 		14: PF - parity flag
	// 		13: SF - Sign flag
	// 		12: IF - Interrupt flag
	// 		11: ZF - Zero flag
	// 		10: OF - Overflow flag
	// 		9: INF - Input flag
	// 		8: OUF - Output flag
	// 		7:0 - unused
	flagRegister []bool // flag register
}

const stackBufferAddr = 0x03FF // 1023

// Starting and ending addresses for the different sections of memory
/* NOTE: These are not the actual starting and ending addresses, the
addresses swapped for the emulator. The stack on the hardware grows down, but
the stack on the emulator grows up.
*/
const wspStartAddr, wspEndAddr = 0x02D0, 0x03FE         // 720, 1022
const dspStartAddr, dspEndAddr = 0x0250, 0x02CF         // 592, 719
const rspStartAddr, rspEndAddr = 0x0240, 0x024F         // 576, 591
const handlerStartAddr, handlerEndAddr = 0x0100, 0x023F // 256, 575

// NewProcessor creates a new processor
func NewProcessor() *Processor {
	return &Processor{
		wsp:            wspStartAddr,
		dsp:            dspStartAddr,
		rsp:            rspStartAddr,
		flagRegister:   make([]bool, 16),
		programCounter: 0x0,
	}
}

// LoadProgram loads a program into memory
func (p *Processor) LoadProgram(program []int16) {
	if len(program) > 255 {
		panic("program exceeds 255 words, the assembler should have caught this!")
	}
	// for i, instruction := range program {
	// 	p.ram[i] = instruction
	// }
	copy(program, p.ram[0:len(program)])
}

// // imm is a 10 bit signed integer, so we need to sign extend it
// func (p *Processor) signExtImm(imm int16) int16 {
// 	if imm&0x200 != 0 {
// 		imm |= 0xFC00
// 	}
// 	return imm
// }

// func (p *Processor) getNextOpcode() int16 {
// 	// the opcdode is the first 6 bits of the instruction
// 	return p.ram[p.programCounter] >> 10
// }

func (p *Processor) DecodeInstruction() {
	instruction := p.ram[p.programCounter]
	opcode := instruction >> 10
	imm := instruction & 0x03FF
	switch opcode {
	case 0b100000:
		p.add()
	case 0b100001:
		p.addi(imm)
	case 0b100010:
		p.and()
	case 0b100100:
		p.addi(imm)
	case 0b101000:
		p.blt(imm)
	case 0b110000:
		p.bgt(imm)
	case 0b110001:
		p.beq(imm)
	case 0b110010:
		p.call(imm)
	case 0b111010:
		p.cmp()
	case 0b110111:
		p.clr(imm)
	case 0b000000:
		p.exit()
	case 0b101010:
		p.j(imm)
	case 0b101110:
		p.js()
	case 0b101111:
		p.ld(imm)
	case 0b100101:
		p.lui(imm)
	case 0b100110:
		p.mdsp(imm)
	case 0b000001:
		p.ori(imm)
	case 0b000010:
		p.or()
	case 0b000100:
		p.pop(imm)
	case 0b000110:
		p.pushi(imm)
	case 0b001000:
		p.push(imm)
	case 0b001010:
		p.ret()
	case 0b001110:
		p.sfl(imm)
	case 0b010000:
		p.sfr(imm)
	case 0b010010:
		p.st(imm)
	case 0b010011:
		p.sub()
	case 0b011000:
		p.swap()
	default:
		panic("invalid opcode")
	}
}
