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
	commands := strings.Split(string(input), ",")
	programs := make([]string, 0)
	for i := 0; i < 16; i++ {
		programs = append(programs, string(i+97))
	}
	initial := strings.Join(programs, "")
	iterations := 0
	for true {
		programs = dance(programs, commands)
		iterations++
		// search for iteration where string is repeating
		if part == 1 || strings.Join(programs, "") == initial {
			break
		}
	}
	if part == 2 {
		cycle := 1000000000 % iterations
		for i := 0; i < cycle; i++ {
			programs = dance(programs, commands)
		}
	}

	fmt.Println(strings.Join(programs, ""))
}

func dance(programs []string, commands []string) []string {
	for _, cmd := range commands {
		switch cmd[0] {
		case 's':
			arg, _ := strconv.Atoi(cmd[1:])
			programs = append(programs[len(programs)-arg:], programs[:len(programs)-arg]...)
		case 'x':
			args := strings.Split(cmd[1:], "/")
			i, _ := strconv.Atoi(args[0])
			j, _ := strconv.Atoi(args[1])
			programs[i], programs[j] = programs[j], programs[i]
		case 'p':
			var i, j int
			args := strings.Split(cmd[1:], "/")
			for k, v := range programs {
				if v == args[0] {
					i = k
				} else if v == args[1] {
					j = k
				}
			}
			programs[i], programs[j] = programs[j], programs[i]
		}
	}
	return programs
}
