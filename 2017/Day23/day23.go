package main

import (
	"fmt"
	"io/ioutil"
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
	input, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(input), "\r\n")
	registers := make(map[string]int)
	if part == 2 {
		registers["a"] = 1
	}
	count := 0
	for i := 0; i < len(lines); i++ {
		cmd := strings.Split(lines[i], " ")
		switch cmd[0] {
		case "set":
			val, err := strconv.Atoi(cmd[2])
			if err != nil {
				registers[cmd[1]] = registers[cmd[2]]
			} else {
				registers[cmd[1]] = val
			}
			break
		case "sub":
			val, err := strconv.Atoi(cmd[2])
			if err != nil {
				registers[cmd[1]] -= registers[cmd[2]]
			} else {
				registers[cmd[1]] -= val
			}
			break
		case "mul":
			count++
			val, err := strconv.Atoi(cmd[2])
			if err != nil {
				registers[cmd[1]] *= registers[cmd[2]]
			} else {
				registers[cmd[1]] *= val
			}
			break
		case "jnz":
			var test int
			arg, err := strconv.Atoi(cmd[1])
			if err != nil {
				test = registers[cmd[1]]
			} else {
				test = arg
			}
			if test != 0 {
				val, _ := strconv.Atoi(cmd[2])
				i = (i - 1) + val
			}
			break
		}
	}
	if part == 1 {
		fmt.Println(count)
	} else {
		fmt.Println(registers["h"])
	}
}
