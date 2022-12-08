package six

import (
	"fmt"
	"testing"
)

func TestSix(t *testing.T) {
	fmt.Println("Part One", PartOne(ReadInput()))
	fmt.Println("Part Two", PartTwo(ReadInput()))
}
