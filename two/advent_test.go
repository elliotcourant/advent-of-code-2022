package two

import (
	"fmt"
	"testing"
)

func TestTotalScore(t *testing.T) {
	fmt.Println("Day 2 Total Score:", TotalScore(ParseInput()))
	fmt.Println("Day 2 Part 2 Total:", TotalScore(ParseInputPartTwo()))
}
