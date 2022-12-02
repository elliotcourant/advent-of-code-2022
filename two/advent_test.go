package two

import (
	"fmt"
	"testing"
)

func TestTotalScore(t *testing.T) {
	fmt.Println("Day 2 Total Score:", TotalScore(ParseInput()))
	fmt.Println("Day 2 Part 2 Total:", TotalScore(ParseInputPartTwo()))
}

func TestFunk(t *testing.T) {
	fmt.Println("Funk:", Result(Rock, Rock))         // 3 Draw
	fmt.Println("Funk:", Result(Rock, Paper))        // 6 Win
	fmt.Println("Funk:", Result(Rock, Scissors))     // 0 Lose
	fmt.Println("Funk:", Result(Paper, Rock))        // 0
	fmt.Println("Funk:", Result(Paper, Paper))       // 3
	fmt.Println("Funk:", Result(Paper, Scissors))    // 6
	fmt.Println("Funk:", Result(Scissors, Rock))     // 6
	fmt.Println("Funk:", Result(Scissors, Paper))    // 0
	fmt.Println("Funk:", Result(Scissors, Scissors)) // 3
}
