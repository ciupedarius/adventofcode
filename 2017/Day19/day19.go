package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	var part int
	// part is defined as cmd argument
	if len(os.Args) > 1 && os.Args[1] == "part2" {
		part = 2
	} else {
		//run part 1 as default
		part = 1
	}
	input, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(input), "\r\n")
	diagram := make([][]string, len(lines))
	for i := 0; i < len(lines); i++ {
		diagram[i] = make([]string, len(lines[i]))
		for j, c := range lines[i] {
			diagram[i][j] = string(c)
		}
	}
	path := make([]string, 0)
	steps := 1
	direction := 2 // 0=up, 2=down, 4=left, 6=right
	// find starting point for diagram
	for i, v := range diagram[0] {
		if v == "|" {
			traverseDiagram(diagram, 1, i, direction, &path, &steps)
			if part == 1 {
				fmt.Println(strings.Join(path, ""))
			} else {
				fmt.Println(steps)
			}
		}
	}

}

func traverseDiagram(diagram [][]string, i int, j int, direction int, path *[]string, steps *int) {
	isEnd := false
	*steps++
	switch diagram[i][j] {
	case "|":
		if direction < 4 {
			i = i - (1 - direction)
		} else {
			j = j - (5 - direction)
		}
		break
	case "-":
		if direction < 4 {
			i = i - (1 - direction)
		} else {
			j = j - (5 - direction)
		}

		break
	case "+":
		// get all neighbors for '+'
		next := []string{diagram[i-1][j], diagram[i+1][j], diagram[i][j-1], diagram[i][j+1]}
		for k, v := range next {
			var inverseDirection int
			if (direction/2)%2 == 0 {
				inverseDirection = (direction / 2) + 1
			} else {
				inverseDirection = (direction / 2) - 1
			}
			// check if new path is different from the one it came
			if v != " " && k != inverseDirection {
				direction = k * 2
				break
			}
		}
		if direction < 4 {
			i = i - (1 - direction)
		} else {
			j = j - (5 - direction)
		}
		break
	case " ":
		break
	// letters
	default:
		*path = append(*path, diagram[i][j])
		if diagram[i][j] == "Z" {
			isEnd = true
		}
		if !isEnd {
			if direction < 4 {
				i = i - (1 - direction)
			} else {
				j = j - (5 - direction)
			}
		}
	}
	if isEnd {
		return
	} else {
		traverseDiagram(diagram, i, j, direction, path, steps)
	}
}
