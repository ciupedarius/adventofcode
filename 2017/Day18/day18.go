package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"sync"
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
	var wg sync.WaitGroup
	input, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(input), "\r\n")
	if part == 1 {
		registers := make(map[string]int)
		lastFrequency := execInstructions(lines, registers, part, &wg)
		fmt.Println(lastFrequency)
	} else {
		firstProgram := make(map[string]int)
		secondProgram := make(map[string]int)
		secondProgram["p"] = 1
		wg.Add(1)
		go execInstructions(lines, firstProgram, part, &wg)
		wg.Add(1)
		go execInstructions(lines, secondProgram, part, &wg)
		wg.Wait()
	}
}

func execInstructions(lines []string, registers map[string]int, part int, wg *sync.WaitGroup) int {
	received := false //part 1
	result := 0
	for i := 0; i < len(lines); i++ {
		row := strings.Split(lines[i], " ")
		switch row[0] {
		case "set":
			value, err := strconv.Atoi(row[2])
			if err != nil {
				value = registers[row[2]]
			}
			registers[row[1]] = value
			break
		case "snd":
			value, err := strconv.Atoi(row[1])
			if err != nil {
				value = registers[row[1]]
			}
			if part == 1 {
				result = value
			}
			break
		case "add":
			value, err := strconv.Atoi(row[2])
			if err != nil {
				value = registers[row[2]]
			}
			registers[row[1]] += value
			break
		case "mul":
			value, err := strconv.Atoi(row[2])
			if err != nil {
				value = registers[row[2]]
			}
			registers[row[1]] *= value
			break
		case "mod":
			value, err := strconv.Atoi(row[2])
			if err != nil {
				value = registers[row[2]]
			}
			registers[row[1]] = registers[row[1]] % value
			break
		case "rcv":
			if part == 1 {
				if result != 0 {
					received = true
				}
			}
			break
		case "jgz":
			if registers[row[1]] > 0 {
				value, err := strconv.Atoi(row[2])
				if err != nil {
					value = registers[row[2]]
				}
				i = i - 1 + value
			}
			break
		}
		if received {
			break
		}
	}
	if part == 2 {
		wg.Done()
	}
	return result
}
