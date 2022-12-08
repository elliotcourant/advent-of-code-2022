package four

import (
	"fmt"
	"testing"
)

func TestRange(t *testing.T) {
	fmt.Println("Part One", PartOne(ReadInput()))
	fmt.Println("Part Two", PartTwo(ReadInput()))
}
