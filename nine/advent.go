package nine

import "embed"

//go:embed *.txt
var input embed.FS

func ParseInput() {
	data, err := input.ReadFile("input.txt")
	if err != nil {
		panic("Failed to read the input data!")
	}
}
