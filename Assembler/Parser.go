package assembler

import (
	"bufio"
	"errors"
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

func (p *Parser) Parse() error {
	labels := make(map[string]int)
	if len(p.lines) > 256 {
		return errors.New("file exceeds 256 lines")
	}

	// First pass, removes comments, empty lines. Removes labels and adds them to the map
	for i := len(p.lines) - 1; i >= 0; i-- {
		// Checks for comments and removes them
		if strings.Contains(p.lines[i], "#") {
			p.lines[i] = p.lines[i][:strings.Index(p.lines[i], "#")]
		}
		p.lines[i] = strings.TrimSpace(p.lines[i])
		// Removes empty lines
		if p.lines[i] == "" {
			p.lines = append(p.lines[:i], p.lines[i+1:]...)
		}
	}

	// Second pass, removes labels and adds them to the map
	for i := 0; i < len(p.lines); i++ {
		if strings.Contains(p.lines[i], ":") {
			labels[p.lines[i][:strings.Index(p.lines[i], ":")]] = i
			// - p.instOffset
			// p.instOffset++
			// p.lines[i] = p.lines[i][strings.Index(p.lines[i], ":")+1:]
			// remove the line
			p.lines = append(p.lines[:i], p.lines[i+1:]...)
		}
	}

	// Third pass, replaces label addresses with their actual addresses
	for i, line := range p.lines {
		for label, address := range labels {
			if strings.Contains(line, label) {
				if strings.Contains(line, "beq") || strings.Contains(line, "bgt") {
					newAddress := address - i - 1
					p.lines[i] = strings.Replace(line, label, fmt.Sprint(newAddress), 1)
				} else {
					p.lines[i] = strings.Replace(line, label, fmt.Sprint(address), 1)
				}
			}
		}
	}
	return nil
}
