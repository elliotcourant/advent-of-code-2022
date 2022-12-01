package one

import (
	"fmt"
	"testing"
)

func TestWhoHasTheMost(t *testing.T) {
	input := ReadInput()

	_, most := WhoHasTheMost(input)
	fmt.Println("Total:", most, "has the most")
}

func TestTopThree(t *testing.T) {
	input := ReadInput()

	fmt.Println("Total:", TopN(input, 3), "are the total  3")
}
