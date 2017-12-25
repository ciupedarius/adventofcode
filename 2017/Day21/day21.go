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
	rules := make(map[string]string)
	patterns := []string{".#./..#/###"}
	for _, line := range lines {
		l := strings.Split(line, "=>")
		rules[strings.TrimSpace(l[0])] = strings.TrimSpace(l[1])
	}
	inputs := make([]string, 0)
	iterations := 5
	if part == 2 {
		iterations = 18
	}
	for k := 0; k < iterations; k++ {
		newPatterns := make([]string, 0) // pattern after 1 iteration
		for _, pattern := range patterns {
			grids := make([][][]string, 0)
			size := 0 // find if it's a 3x3 or a 4x4 grid
			for _, v := range pattern {
				if v == '/' {
					size++
				}
			}
			if size%2 == 0 { // 3x3 grid
				grid := patternToGrid(pattern)
				grids = append(grids, grid)
			} else { //need to split grid
				for _, split := range splitPattern(pattern) {
					grid := patternToGrid(split)
					grids = append(grids, grid)
				}
			}
			newGridPatterns := make([]string, 0)
			for _, grid := range grids {
				var newGridPattern string
				found := false
				for i := 0; i < 4; i++ {
					grid = rotateGrid(grid)
					pattern = gridToPattern(grid)
					inputs = append(inputs, pattern)
					newGridPattern, found = getPattern(pattern, rules)
					if found {
						break
					} else {
						flipped := flipPattern(pattern)
						inputs = append(inputs, flipped)
						newGridPattern, found = getPattern(flipped, rules)
						if found {
							break
						}
					}
				}
				if !found {
					panic("No rule found for grid")
				}
				newGridPatterns = append(newGridPatterns, newGridPattern)
			}
			if len(newGridPatterns) == 1 {
				newPatterns = append(newPatterns, newGridPatterns...)
			} else if len(newGridPatterns) == 9 {
				newPatterns = append(newPatterns, newGridPatterns...)
			} else {
				newPatterns = append(newPatterns, concatenatePatterns(newGridPatterns))
			}
		}
		patterns = newPatterns
	}
	count := 0
	for _, p := range patterns {
		for _, c := range p {
			if c == '#' {
				count++
			}
		}
	}
	fmt.Println(count)
}

func concatenatePatterns(patterns []string) string {
	pattern := ""
	lastChar := "/"
	for i := 0; i < len(patterns); i += 2 {
		if i == len(patterns)-2 {
			lastChar = ""
		}
		s1 := strings.Split(patterns[i], "/")
		s2 := strings.Split(patterns[i+1], "/")
		pattern = pattern + s1[0] + s2[0] + "/" + s1[1] + s2[1] + "/" + s1[2] + s2[2] + lastChar
	}
	return pattern
}

func getPattern(pattern string, rules map[string]string) (string, bool) {
	v, ok := rules[pattern]
	if !ok {
		return "", false
	} else {
		return v, true
	}
}

func patternToGrid(pattern string) [][]string {
	lines := strings.Split(pattern, "/")
	grid := make([][]string, len(lines))
	for i := 0; i < len(lines); i++ {
		grid[i] = make([]string, len(lines))
		for j, c := range lines[i] {
			grid[i][j] = string(c)
		}
	}
	return grid
}

func gridToPattern(grid [][]string) string {
	pattern := ""
	lastChar := "/"
	for i := 0; i < len(grid[0]); i++ {
		if i == len(grid[0])-1 {
			lastChar = ""
		}
		pattern = pattern + strings.Join(grid[i], "") + lastChar
	}
	return pattern
}

func flipPattern(pattern string) string {
	arr := []byte(pattern)
	if len(arr) > 5 {
		arr[0], arr[2] = arr[2], arr[0]
		arr[4], arr[6] = arr[6], arr[4]
		arr[8], arr[10] = arr[10], arr[8]
	} else {
		arr[0], arr[1] = arr[1], arr[0]
		arr[3], arr[4] = arr[4], arr[3]
	}
	return string(arr)
}

func rotateGrid(grid [][]string) [][]string {
	rotated := make([][]string, len(grid[0]))
	for i := 0; i < len(grid[0]); i++ {
		rotated[i] = make([]string, len(grid[0]))
		for j := 0; j < len(grid[0]); j++ {
			rotated[i][j] = grid[len(grid[0])-j-1][i]
		}
	}
	return rotated
}

func splitPattern(pattern string) []string {
	lines := strings.Split(pattern, "/")
	patterns := make([]string, 0)
	if len(lines) == 4 {
		for i := 0; i < 3; i += 2 {
			line1 := []byte(lines[i])
			line2 := []byte(lines[i+1])
			newPattern := []byte{line1[0], line1[1], '/', line2[0], line2[1]}
			patterns = append(patterns, string(newPattern))
			newPattern = []byte{line1[2], line1[3], '/', line2[2], line2[3]}
			patterns = append(patterns, string(newPattern))
		}
	} else {
		for i := 0; i < 6; i += 2 {
			line1 := []byte(lines[i])
			line2 := []byte(lines[i+1])
			newPattern := []byte{line1[0], line1[1], '/', line2[0], line2[1]}
			patterns = append(patterns, string(newPattern))
			newPattern = []byte{line1[2], line1[3], '/', line2[2], line2[3]}
			patterns = append(patterns, string(newPattern))
			newPattern = []byte{line1[4], line1[5], '/', line2[4], line2[5]}
			patterns = append(patterns, string(newPattern))
		}
	}
	return patterns
}
