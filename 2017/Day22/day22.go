package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type node struct {
	x, y int
}

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
	grid := make(map[node]int)

	for i, line := range lines {
		for j, c := range line {
			if c == '.' {
				grid[node{j, i}] = 0 // 0=clean, 1=weakened, 2=infected, 3=flagged
			} else {
				grid[node{j, i}] = 2
			}
		}
	}
	y := len(lines) / 2
	x := len(lines[0]) / 2
	if part == 1 {
		part1(grid, x, y)
	} else {
		part2(grid, x, y)
	}
}

func part1(grid map[node]int, x int, y int) {
	count := 0
	direction := 1 //0=left 1=up 2=right 3=down
	for i := 0; i < 10000; i++ {
		val, ok := grid[node{x, y}]
		if !ok {
			// unvisited, clean node
			direction-- //turn to left
			if direction < 0 {
				direction = 3
			}
			grid[node{x, y}] = 2
			count++
			x, y = moveNode(direction, x, y)
		} else { // node in grid
			if val == 2 { //infected node
				direction++ // move to right
				direction = direction % 4
				grid[node{x, y}] = 0
				x, y = moveNode(direction, x, y)
			} else {
				direction-- //turn to left
				if direction < 0 {
					direction = 3
				}
				grid[node{x, y}] = 2
				count++
				x, y = moveNode(direction, x, y)
			}
		}
	}
	fmt.Println(count)
}

func part2(grid map[node]int, x int, y int) {
	count := 0
	direction := 1 //0=left 1=up 2=right 3=down
	for i := 0; i < 10000000; i++ {
		val, ok := grid[node{x, y}]
		if !ok {
			// unvisited, clean node
			direction-- //turn to left
			if direction < 0 {
				direction = 3
			}
			grid[node{x, y}] = 1
			x, y = moveNode(direction, x, y)
		} else { // node in grid
			switch val {
			case 0: //clean
				direction-- //turn to left
				if direction < 0 {
					direction = 3
				}
				grid[node{x, y}] = 1
				x, y = moveNode(direction, x, y)
				break
			case 1: // weakened
				grid[node{x, y}] = 2
				count++
				x, y = moveNode(direction, x, y)
				break
			case 2: //infected
				direction++ // move to right
				direction = direction % 4
				grid[node{x, y}] = 3
				x, y = moveNode(direction, x, y)
				break
			case 3: //flagged
				direction += 2 // reverse direction
				direction = direction % 4
				grid[node{x, y}] = 0
				x, y = moveNode(direction, x, y)
				break
			}
		}
	}
	fmt.Println(count)
}

func moveNode(direction int, x int, y int) (int, int) {
	switch direction {
	case 0:
		x--
		break
	case 1:
		y--
		break
	case 2:
		x++
		break
	case 3:
		y++
		break
	}
	return x, y
}
