package four

import (
	"bytes"
	"embed"
	"strconv"
)

//go:embed *.txt
var input embed.FS

type Elf []int

func (e Elf) Contains(other Elf) bool {
	myFirst, myLast := e[0], e[len(e)-1]
	otherFirst, otherLast := other[0], other[len(other)-1]
	return otherFirst >= myFirst && otherLast <= myLast
}

func (e Elf) Overlap(other Elf) bool {
	myFirst, myLast := e[0], e[len(e)-1]
	otherFirst, otherLast := other[0], other[len(other)-1]
	return otherLast >= myFirst && otherFirst <= myLast
}

type Pair [2]Elf

func (p Pair) FullyContains() bool {
	return p[0].Contains(p[1]) || p[1].Contains(p[0])
}

func (p Pair) Overlap() bool {
	return p[0].Overlap(p[1])
}

func ReadInput() []Pair {
	data, err := input.ReadFile("input.txt")
	if err != nil {
		panic("Failed to read the input data!")
	}

	lines := bytes.Split(data, []byte("\n"))
	result := make([]Pair, 0, len(lines))
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		parts := bytes.SplitN(line, []byte(","), 2)
		var pair Pair
		for i := 0; i < len(parts); i++ {
			pair[i] = ParsePart(parts[i])
		}
		result = append(result, pair)
	}
	return result
}

func ParsePart(part []byte) Elf {
	pieces := bytes.SplitN(part, []byte("-"), 2)
	start, err := strconv.ParseInt(string(pieces[0]), 10, 64)
	if err != nil {
		panic(err)
	}
	end, err := strconv.ParseInt(string(pieces[1]), 10, 64)
	if err != nil {
		panic(err)
	}
	return NewElf(int(start), int(end), true)
}

func NewElf(start int, end int, inc bool) Elf {
	offset := 0
	if inc {
		offset += 1
	}
	result := make([]int, (end-start)+offset)
	for i := 0; i < len(result); i++ {
		result[i] = start + i
	}
	return result
}

func PartOne(pairs []Pair) int {
	var count int
	for _, pair := range pairs {
		if pair.FullyContains() {
			count++
		}
	}

	return count
}

func PartTwo(pairs []Pair) int {
	var count int
	for _, pair := range pairs {
		if pair.Overlap() {
			count++
		}
	}

	return count
}
