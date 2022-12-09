package eight

import "embed"

//go:embed *.txt
var input embed.FS

type Forest [][]int

func ReadInput() {
	data, err := input.ReadFile("input.txt")
	if err != nil {
		panic("Failed to read the input data!")
	}
}
