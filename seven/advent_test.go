package seven

import (
	"fmt"
	"testing"
)

func TestSeven(t *testing.T) {
	fmt.Println("Part One", PartOne(BuildFileSystem(ReadInput())))
	fmt.Println("Part Two", PartTwo(BuildFileSystem(ReadInput())))
}
