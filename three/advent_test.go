package three

import (
	"fmt"
	"testing"
)

func TestDayThree(t *testing.T) {
	fmt.Println(PartOne(ReadInput()))
	fmt.Println(PartTwo(ReadInputPartTwo()))
}
