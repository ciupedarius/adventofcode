package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	inFile, _ := os.Open("input.txt")
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)
	set := make(map[string]int)
	max := 0
	for scanner.Scan() {
		row := strings.Fields(scanner.Text())
		result := executeOp(row[4], row[6], row[5], set)
		if result {
			set[row[0]] = executeAction(set, row[0], row[2], row[1])
		}
		// for part 2 compute max on each iteration
		if part == 2 {
			max = getMax(set, max)
		}
	}
	// compute only last max for part 1
	if part == 1 {
		max = getMax(set, max)
	}
	fmt.Println(max)
}

func executeOp(a string, b string, op string, set map[string]int) bool {
	aValue := set[a]
	bValue, _ := strconv.Atoi(b)
	switch op {
	case "<":
		return aValue < bValue
	case "<=":
		return aValue <= bValue
	case ">":
		return aValue > bValue
	case ">=":
		return aValue >= bValue
	case "==":
		return aValue == bValue
	case "!=":
		return aValue != bValue
	default:
		fmt.Println("Unsupported op:", op)
		return false
	}
}

func executeAction(set map[string]int, a string, b string, act string) int {
	val, _ := strconv.Atoi(b)
	if act == "inc" {
		return set[a] + val
	} else {
		return set[a] - val
	}
}

func getMax(set map[string]int, crtMax int) int {
	max := 0
	for _, v := range set {
		if v > max {
			max = v
		}
	}
	if max > crtMax {
		return max
	}
	return crtMax
}
