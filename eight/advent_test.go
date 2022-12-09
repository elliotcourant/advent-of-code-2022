package eight

import (
	"fmt"
	"testing"
)

func TestDayEight(t *testing.T) {
	forest := ReadInput()
	fmt.Println("Part One", PartOne(forest))
	fmt.Println("Part Two", PartTwo(forest))
}
