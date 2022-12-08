package six

import (
	"bytes"
	"embed"
)

//go:embed *.txt
var input embed.FS

func ReadInput() []byte {
	data, err := input.ReadFile("input.txt")
	if err != nil {
		panic("Failed to read the input data!")
	}

	return data
}

func Start(data []byte, size int) int {
OuterLoop:
	for i := size; i < len(data); i++ {
		chunk := data[i-size : i]
		// Read the chunk going forward.
		for a, b := range chunk {
			// But find the current byte reading backward. If we find a byte without the same index
			// then that means we have found a duplicate byte. This isnt the start of the thing, so
			// just continue the outer loop
			// Is this just a poor mans binary search?
			if bytes.LastIndex(chunk, []byte{b}) != a {
				continue OuterLoop
			}
		}
		return i
	}
	return -1
}

func PartOne(data []byte) int {
	return Start(data, 4)
}

func PartTwo(data []byte) int {
	return Start(data, 14)
}
