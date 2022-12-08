package seven

import (
	"embed"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

//go:embed *.txt
var input embed.FS

type Tree interface {
	Size() int
}

type File int

func (f File) Size() int {
	return int(f)
}

type Directory map[string]Tree

func (d Directory) Size() int {
	var sum int
	for _, item := range d {
		sum += item.Size()
	}

	return sum
}

func ReadInput() []Command {
	data, err := input.ReadFile("input.txt")
	if err != nil {
		panic("Failed to read the input data!")
	}

	expr := regexp.MustCompile(`(?m)^\$\s([^$]+)`)
	commandsAndResults := expr.FindAllStringSubmatch(string(data), -1)
	path := make([]string, 0)
	commands := make([]Command, 0)
	for _, item := range commandsAndResults {
		parts := strings.SplitN(item[1], "\n", 2)
		command := strings.TrimSpace(parts[0])
		results := strings.Split(strings.TrimSpace(parts[1]), "\n")
		// Dumb hack
		if results[0] == "" {
			results = []string{}
		}
		if strings.HasPrefix(command, "cd") {
			if command == "cd .." {
				path = path[:len(path)-1]
			} else if command == "cd /" {
				path = []string{}
			} else {
				path = append(path, strings.TrimPrefix(command, "cd "))
			}
		}
		commands = append(commands, Command{
			Path:    fmt.Sprintf("/%s", strings.Join(path, "/")),
			Command: command,
			Results: results,
		})
	}

	return commands
}

type Command struct {
	Path    string
	Command string
	Results []string
}

func BuildFileSystem(data []Command) *Directory {
	dir := Directory{}
	for _, cmd := range data {
		if strings.HasPrefix(cmd.Command, "cd") {
			continue
		}
		pathParts := strings.Split(strings.TrimPrefix(cmd.Path, "/"), "/")
		cwd := dir
		for _, part := range pathParts {
			if part == "" {
				continue
			}
			nwd, ok := cwd[part]
			if !ok {
				nwd = &Directory{}
				cwd[part] = nwd
			}
			cwd = *(nwd.(*Directory))
		}

		for _, item := range cmd.Results {
			if strings.HasPrefix(item, "dir") {
				cwd[strings.TrimPrefix(item, "dir ")] = &Directory{}
				continue
			}

			fileData := strings.SplitN(item, " ", 2)
			size, err := strconv.ParseInt(fileData[0], 10, 64)
			if err != nil {
				panic(err)
			}

			cwd[fileData[1]] = File(size)
		}
	}
	return &dir
}

func PartOne(filesystem *Directory) int {
	result := FindLessThanOrEqualTo(filesystem, 100000)
	var sum int
	for _, item := range result {
		sum += item.Size()
	}
	return sum
}

func FindLessThanOrEqualTo(filesystem *Directory, max int) []*Directory {
	result := make([]*Directory, 0)
	for _, dir := range *filesystem {
		switch x := dir.(type) {
		case File:
			continue
		case *Directory:
			if dir.Size() <= max {
				result = append(result, x)
			}
			result = append(result, FindLessThanOrEqualTo(x, max)...)
		}
	}

	return result
}

func PartTwo(filesystem *Directory) int {
	available := 70000000 - filesystem.Size()
	deleteAtLeast := 30000000 - available
	return FindSmallestGreaterThan(filesystem, deleteAtLeast)
}

func FindGreaterThan(filesystem *Directory, max int) []*Directory {
	result := make([]*Directory, 0)
	for _, dir := range *filesystem {
		switch x := dir.(type) {
		case File:
			continue
		case *Directory:
			if dir.Size() > max {
				result = append(result, x)
			}
			result = append(result, FindGreaterThan(x, max)...)
		}
	}

	return result
}

func FindSmallestGreaterThan(filesystem *Directory, max int) int {
	result := FindGreaterThan(filesystem, max)
	// The dreaded sort. :rip:
	sort.Slice(result, func(i, j int) bool {
		return result[i].Size() < result[j].Size()
	})

	return result[0].Size()
}
