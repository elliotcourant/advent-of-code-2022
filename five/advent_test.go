package five

import (
	"fmt"
	"testing"
)

func TestFive(t *testing.T) {
	layout, instructions := ReadInput()

	fmt.Println(PartOne(layout, instructions))

	layout, instructions = ReadInput()
	fmt.Println(PartTwo(layout, instructions))
}
