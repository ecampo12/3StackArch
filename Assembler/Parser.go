package assembler

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Parser struct {
	fileName   string
	wordLenght int
	instOffset int
	lines      []string
}

func NewParser(fileName string) *Parser {
	var p Parser = Parser{
		fileName:   fileName,
		wordLenght: 1,
		instOffset: 0,
	}
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
		p.lines = append(p.lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return &p
}

func (p *Parser) GetLines() []string {
	return p.lines
}

func (p *Parser) Parse() {
	labels := make(map[string]int)
	if len(p.lines) > 256 {
		return // TODO: throw error, too many lines
	}

	// First pass, removes comments and empty lines
	for i := len(p.lines) - 1; i >= 0; i-- {
		p.lines[i] = strings.TrimSpace(p.lines[i])
		// Checks for comments and removes them
		if strings.Contains(p.lines[i], "#") {
			p.lines[i] = p.lines[i][:strings.Index(p.lines[i], "#")]
		}
		// Removes empty lines
		if p.lines[i] == "" {
			p.lines = append(p.lines[:i], p.lines[i+1:]...)
		}
	}

	// Second pass, removes labels and adds them to the map
	for i, line := range p.lines {
		if strings.Contains(line, ":") {
			labels[line[:strings.Index(line, ":")]] = i - p.instOffset
			p.instOffset++
			p.lines[i] = line[strings.Index(line, ":")+1:]
		}
	}

	// Third pass, replaces label addresses with their actual addresses
	for i, line := range p.lines {
		for label, address := range labels {
			if strings.Contains(line, label) {
				if strings.Contains(line, "beq") || strings.Contains(line, "bgt") {
					newAddress := address - (i + 1)
					p.lines[i] = strings.Replace(line, label, fmt.Sprint(newAddress), 1)
				} else {
					p.lines[i] = strings.Replace(line, label, fmt.Sprint(address), 1)
				}
			}
		}
	}
}
