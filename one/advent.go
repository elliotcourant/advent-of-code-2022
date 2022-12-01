package one

import (
	"bytes"
	"embed"
	"strconv"
)

//go:embed *.txt
var input embed.FS

func ReadInput() map[int][]int {
	data, err := input.ReadFile("input.txt")
	if err != nil {
		panic("Failed to read the input data!")
	}

	result := map[int][]int{
		1: make([]int, 0), // Seed the first one
	}
	for _, item := range bytes.Split(data, []byte("\n")) {
		if len(item) == 0 {
			// New elf
			result[len(result)+1] = make([]int, 0)
			continue
		}

		elfNumber := len(result)
		number, err := strconv.ParseInt(string(item), 10, 64)
		if err != nil {
			panic(err)
		}

		result[elfNumber] = append(result[elfNumber], int(number))
	}

	return result
}

func WhoHasTheMost(calories map[int][]int) (elf int, total int) {
	var number, most int
	for elf, items := range calories {
		var total int
		for _, item := range items {
			total += item
		}
		if total > most {
			number = elf
			most = total
		}
	}

	return number, most
}

func TopN(calories map[int][]int, n int) int {
	data := calories
	var total int
	for i := 0; i < n; i++ {
		elf, elfsCalories := WhoHasTheMost(data)
		total += elfsCalories
		// Take the current "top elf" out of the dataset, and do it again
		delete(data, elf)
	}

	return total
}
