package assembler

import "fmt"

func main() {
	p := NewParser("/test.asm")
	lines := p.GetLines()
	for _, line := range lines {
		fmt.Println(line)
	}
}
