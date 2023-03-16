This is a go implementation of the assembler and emulator for the Three Stack Architecture, a processor my group and I designed for our Computer Architecture course at [RHIT] (https://www.rose-hulman.edu/).

For details, please see the [project report] (Design Documents/Design Document.pdf).

## Assembler
The assembler is a command line tool that takes a file containing Three Stack Assembly code and outputs a binary file that can be loaded into the emulator.

### Usage
```
$ go build assembler.go
$ ./assembler <input file> <output file>
```

## Emulator
The emulator is a command line tool that takes a binary file containing Three Stack Assembly code and executes it.
