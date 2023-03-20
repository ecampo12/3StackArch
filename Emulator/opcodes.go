package Emulator

import (
	"fmt"
)

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

// M[$wsp] = M[$wsp+2] + M[$wsp]
// $wsp = $wsp + 2
func (p *Processor) add() {
	p.ram[p.wsp+1] = p.ram[p.wsp+1] + p.ram[p.wsp]
	p.wsp += 1
}

// M[$wsp] = M[$wsp] + SignExtImm
func (p *Processor) addi(imm uint16) {
	p.ram[p.wsp] = p.ram[p.wsp] + imm
}

// M[$wsp] = M[$wsp+2] & M[$wsp]
// $wsp = $wsp + 2
func (p *Processor) and() {
	p.ram[p.wsp] = p.ram[p.wsp+1] & p.ram[p.wsp]
	p.wsp += 1
}

// M[$wsp] = M[$wsp] & ZeroExtImm
func (p *Processor) andi(imm uint16) {
	p.ram[p.wsp] = p.ram[p.wsp] & imm
}

// (SF == 1) ? $pc = $pc + 2 + BranchAddr :
// $pc = $pc +2
func (p *Processor) blt(imm uint16) {
	if p.flagRegister[13] {
		p.programCounter = p.programCounter + 1 + imm
	} else {
		p.programCounter += 1
	}
}

// (SF == 0 && ZF == 0) ? $pc = $pc + 2 +BranchAddr : $pc = $pc +2
func (p *Processor) bgt(imm uint16) {
	if !p.flagRegister[13] && !p.flagRegister[11] {
		p.programCounter = p.programCounter + 1 + imm
	} else {
		p.programCounter += 1
	}
}

// (ZF == 1) ? $pc = $pc + 2 + BranchAddr : $pc = $pc +2
func (p *Processor) beq(imm uint16) {
	if p.flagRegister[11] {
		p.programCounter = p.programCounter + 1 + imm
	} else {
		p.programCounter += 1
	}
}

// $rsp = $rsp – 2 , M[$rsp] = $pc + 2, $pc = CalleeAddr
func (p *Processor) call(imm uint16) {
	p.rsp -= 1
	p.ram[p.rsp] = p.programCounter + 1
	p.programCounter = imm
}

// if M[$wsp + 2] > M[$wsp] : SF = 0, ZF = 0
// else if M[$wsp + 2] == M[$wsp]: ZF = 1
// else if M[$wsp + 2] < M[$wsp]: SF = 1, ZF = 0
// $wsp = $wsp + 4
func (p *Processor) cmp() {
	if p.ram[p.wsp+1] > p.ram[p.wsp] {
		p.flagRegister[13] = false
		p.flagRegister[11] = false
	} else if p.ram[p.wsp+1] == p.ram[p.wsp] {
		p.flagRegister[11] = true
	} else if p.ram[p.wsp+1] < p.ram[p.wsp] {
		p.flagRegister[13] = true
		p.flagRegister[11] = false
	}
	p.wsp += 4
}

// $wsp = $wsp + (SignExtImm<<1)
func (p *Processor) clr(imm uint16) {
	p.wsp += imm << 1
}

func (p *Processor) exit() {
	fmt.Println("Program exited successfully")
	fmt.Println("Top of working stack: ", p.ram[p.wsp])
	// os.Exit(0)
	p.programExit = true
}

// $pc = JumpAddr
func (p *Processor) j(imm uint16) {
	p.programCounter = imm
}

// $pc = M[$wsp], $wsp = $wsp + 2
func (p *Processor) js() {
	p.programCounter = p.ram[p.wsp]
	p.wsp += 1
}

// $wsp = $wsp – 2, M[$wsp] = M[ $dsp + SignExtImm<<1]
func (p *Processor) ld(imm uint16) {
	p.wsp -= 1
	p.ram[p.wsp] = p.ram[p.dsp+(imm<<1)]
}

// $wsp = $wsp – 2, M[$wsp] = Upper8bit(Imm) 8’b0
func (p *Processor) lui(imm uint16) {
	p.wsp -= 1
	p.ram[p.wsp] = imm << 8
}

// $dsp = $dsp + Imm<<1
func (p *Processor) mdsp(imm uint16) {
	p.dsp += imm << 1
}

// M[$wsp] = M[$wsp] | ZeroExtImm
func (p *Processor) ori(imm uint16) {
	p.ram[p.wsp] = p.ram[p.wsp] | imm
}

// M[$wsp] = M[$wsp] | M[$wsp + 2]
// $wsp = $wsp + 2
func (p *Processor) or() {
	p.ram[p.wsp] = p.ram[p.wsp] | p.ram[p.wsp+1]
	p.wsp += 1
}

// M[MemAddress] = M[$wsp], $wsp = $wsp + 2
func (p *Processor) pop(imm uint16) {
	p.ram[imm] = p.ram[p.wsp]
	p.wsp += 1
}

// $wsp = $wsp – 2, M[$wsp] = SignExtImm
func (p *Processor) pushi(imm uint16) {
	p.wsp -= 1
	p.ram[p.wsp] = imm
}

// M[$wsp], $wsp = $wsp – 2, M[$wsp] = M[MemAddress]
func (p *Processor) push(imm uint16) {
	p.ram[p.wsp] = p.ram[imm]
	p.wsp -= 1
}

// $pc = M[$rsp], $rsp = $rsp + 2
func (p *Processor) ret() {
	p.programCounter = p.ram[p.rsp]
	p.rsp += 1
}

// M[$wsp] = M[$wsp] << Imm
func (p *Processor) sfl(imm uint16) {
	p.ram[p.wsp] = p.ram[p.wsp] << imm
}

// M[$wsp] = M[$wsp] >> Imm
func (p *Processor) sfr(imm uint16) {
	p.ram[p.wsp] = p.ram[p.wsp] >> imm
}

// M[$dsp + SignExtImm<<1] = M[$wsp]
// $wsp = $wsp + 2
func (p *Processor) st(imm uint16) {
	p.ram[p.dsp+(imm<<1)] = p.ram[p.wsp]
	p.wsp += 1
}

// M[$wsp] = M[$wsp+2] - M[$wsp]
// $wsp = $wsp + 2
func (p *Processor) sub() {
	p.ram[p.wsp+1] = p.ram[p.wsp+1] - p.ram[p.wsp]
	p.wsp += 1
}

// M[$wsp-2] = M[$wsp], M[$wsp] = M[$wsp + 2], M[$wsp + 2] = M[$wsp - 2]
func (p *Processor) swap() {
	p.ram[p.wsp-1] = p.ram[p.wsp]
	p.ram[p.wsp] = p.ram[p.wsp+1]
	p.ram[p.wsp+1] = p.ram[p.wsp-1]
}

// Might add instructions for Interrupts and Exceptions later
