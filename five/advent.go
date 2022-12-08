package five

import (
	"bytes"
	"embed"
	"fmt"
	"regexp"
	"strconv"
)

//go:embed *.txt
var input embed.FS

func ReadInput() (Orientation, []Instruction) {
	data, err := input.ReadFile("input.txt")
	if err != nil {
		panic("Failed to read the input data!")
	}

	orientation, instructionBytes := ReadInitialOrientation(data)
	instructions := ReadInstructions(instructionBytes)

	return orientation, instructions
}

type Instruction [3]int

func (i Instruction) String() string {
	return fmt.Sprintf("move %d from %d to %d", i[0], i[1], i[2])
}

type Crate byte

type Stack []Crate

type Orientation []Stack

func (o Orientation) Move(instruction Instruction) Orientation {
	n, from, to := instruction[0], instruction[1]-1, instruction[2]-1
	fromStack := o[from]
	fromTaken := fromStack[len(fromStack)-n:]
	fromLeft := fromStack[:len(fromStack)-n]
	o[from] = fromLeft
	for i := len(fromTaken) - 1; i >= 0; i-- {
		o[to] = append(o[to], fromTaken[i])
	}
	return o
}

func (o Orientation) MovePartTwo(instruction Instruction) Orientation {
	n, from, to := instruction[0], instruction[1]-1, instruction[2]-1
	fromStack := o[from]
	fromTaken := fromStack[len(fromStack)-n:]
	fromLeft := fromStack[:len(fromStack)-n]
	o[from] = fromLeft
	o[to] = append(o[to], fromTaken...)
	return o
}

func ReadInitialOrientation(data []byte) (_ Orientation, leftover []byte) {
	parts := bytes.SplitN(data, []byte("\n\n"), 2)
	layoutBytes, leftover := parts[0], parts[1]
	// https://regex101.com/r/JcDyR8/1
	expr := regexp.MustCompile(`(.{3})\s?|$`)
	lines := bytes.Split(layoutBytes, []byte("\n"))

	// Get the number of columns based on the first row.
	columns := expr.FindAll(lines[0], -1)
	orientation := make(Orientation, len(columns))
	for y := len(lines) - 2; y >= 0; y-- {
		crates := expr.FindAll(lines[y], -1)
		for x, crate := range crates {
			if len(bytes.TrimSpace(crate)) == 0 {
				continue
			}
			orientation[x] = append(orientation[x], Crate(crate[1]))
		}
	}

	return orientation, leftover
}

func ReadInstructions(data []byte) []Instruction {
	expr := regexp.MustCompile(`^[^\d]+(\d+)[^\d]+(\d+)[^\d]+(\d+)$`)
	lines := bytes.Split(data, []byte("\n"))
	instructions := make([]Instruction, 0, len(lines))
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		parts := expr.FindSubmatch(line)
		parts = parts[1:]
		var instruction Instruction
		for i := 0; i < 3; i++ {
			n, err := strconv.ParseInt(string(parts[i]), 10, 64)
			if err != nil {
				panic(err)
			}
			instruction[i] = int(n)
		}
		instructions = append(instructions, instruction)
	}

	return instructions
}

func PartOne(layout Orientation, instructions []Instruction) []string {
	for _, instruction := range instructions {
		layout.Move(instruction)
	}

	result := make([]string, len(layout))
	for i, stack := range layout {
		if len(stack) == 0 {
			continue
		}
		result[i] = string(stack[len(stack)-1])
	}

	return result
}

func PartTwo(layout Orientation, instructions []Instruction) []string {
	for _, instruction := range instructions {
		layout.MovePartTwo(instruction)
	}

	result := make([]string, len(layout))
	for i, stack := range layout {
		if len(stack) == 0 {
			continue
		}
		result[i] = string(stack[len(stack)-1])
	}

	return result
}
