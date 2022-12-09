package eight

import (
	"bytes"
	"embed"
	"strconv"
)

//go:embed *.txt
var input embed.FS

type Forest [][]int

func reverse(input []int) []int {
	b := make([]int, len(input))
	copy(b, input)
	if len(b) == 0 {
		return b
	}
	return append(reverse(b[1:]), b[0])
}

func (f Forest) Row(y int) []int {
	return f[y]
}

func (f Forest) Column(x int) []int {
	result := make([]int, len(f[0]))
	for row := 0; row < len(f); row++ {
		result[row] = f[row][x]
	}
	return result
}

func max(data []int) int {
	max := 0
	for _, value := range data {
		if value > max {
			max = value
		}
	}
	return max
}

func distToMax(data []int, ceiling int) int {
	max := 0
	dist := 0
	for i, value := range data {
		if value >= max {
			max = value
		}
		dist = i + 1
		if value >= ceiling {
			return i + 1
		}
	}
	return dist
}

func (f Forest) Visible(y, x int) bool {
	if y == 0 || x == 0 {
		return true
	}
	if y == len(f)-1 || x == len(f[y])-1 {
		return true
	}
	row := f.Row(y)
	height := f[y][x]
	rowLeft := row[:x]
	rowRight := row[len(rowLeft)+1:]
	if height > max(rowLeft) || height > max(rowRight) {
		return true
	}
	column := f.Column(x)
	columnTop := column[:y]
	columnBottom := column[len(columnTop)+1:]
	if height > max(columnTop) || height > max(columnBottom) {
		return true
	}
	return false
}

func (f Forest) ScenicScore(y, x int) int {
	row := f.Row(y)
	col := f.Column(x)
	height := f[y][x]
	up := reverse(col[:y])
	down := col[len(up)+1:]
	left := reverse(row[:x])
	right := row[len(left)+1:]

	upScore := distToMax(up, height)
	downScore := distToMax(down, height)
	leftScore := distToMax(left, height)
	rightScore := distToMax(right, height)
	return upScore * downScore * leftScore * rightScore
}

func ReadInput() Forest {
	data, err := input.ReadFile("input.txt")
	if err != nil {
		panic("Failed to read the input data!")
	}

	lines := bytes.Split(data, []byte("\n"))
	forest := make(Forest, 0, len(lines))
	for y, line := range lines {
		if len(line) == 0 {
			continue
		}
		forest = append(forest, []int{})
		for _, item := range line {
			tree, err := strconv.ParseInt(string(item), 10, 64)
			if err != nil {
				panic(err)
			}
			forest[y] = append(forest[y], int(tree))
		}
	}
	return forest
}

func PartOne(forest Forest) int {
	var total int

	for y := 0; y < len(forest); y++ {
		for x := 0; x < len(forest[y]); x++ {
			if forest.Visible(y, x) {
				total++
			}
		}
	}

	return total
}

func PartTwo(forest Forest) int {
	var max int
	for y := 0; y < len(forest); y++ {
		for x := 0; x < len(forest[y]); x++ {
			if score := forest.ScenicScore(y, x); score > max {
				max = score
			}
		}
	}

	return max
}
