package three

import (
	"bytes"
	"embed"
	"fmt"
)

func init() {
	fmt.Sprint()
}

//go:embed *.txt
var input embed.FS

type Rucksack [2][]byte

func (r Rucksack) AllItems() []byte {
	return append(r[0], r[1]...)
}

func ReadInput() []Rucksack {
	data, err := input.ReadFile("input.txt")
	if err != nil {
		panic("Failed to read the input data!")
	}

	lines := bytes.Split(data, []byte("\n"))
	result := make([]Rucksack, 0, len(lines))
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		half := len(line) / 2
		result = append(result, Rucksack{
			line[:half],
			line[half:],
		})
	}

	return result
}

func ReadInputPartTwo() [][3]Rucksack {
	data := ReadInput()
	numberOfGroups := len(data) / 3
	result := make([][3]Rucksack, numberOfGroups)
	for i, rucksack := range data {
		sack := (i % 3)
		group := (i - sack) % numberOfGroups
		result[group][sack] = rucksack
	}

	return result
}

func PartOne(sacks []Rucksack) int {
	var totalPriority int
	for _, rucksack := range sacks {
		overlap := map[byte]struct{}{}
		Overlap(overlap, rucksack[0], rucksack[1])
		for item := range overlap {
			totalPriority += Priority(item)
		}
	}
	return totalPriority
}

func PartTwo(groups [][3]Rucksack) int {
	var totalPriority int
	for _, group := range groups {
		overlap := map[byte]struct{}{}
		Overlap(overlap, group[0].AllItems(), group[1].AllItems(), group[2].AllItems())
		for item := range overlap {
			totalPriority += Priority(item)
		}
	}
	return totalPriority
}

func Priority(item byte) int {
	if item >= 'a' && item <= 'z' {
		return int(item) - 96
	}
	return int(item) - 38
}

func Overlap(result map[byte]struct{}, a []byte, others ...[]byte) {
BigLoop:
	for _, item := range a {
		for _, other := range others {
			if bytes.IndexByte(other, item) == -1 {
				continue BigLoop
			}
		}
		result[item] = struct{}{}
	}
}
