package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	var part int
	input, _ := ioutil.ReadFile("input.txt")
	// part is defined as cmd argument
	if len(os.Args) > 1 && os.Args[1] == "part2" {
		part = 2
	} else {
		//run part 1 as default
		part = 1
	}

	score := 0
	level := 0
	inGarbage := false
	ignoreNext := false
	characters := 0
	for _, v := range string(input) {
		if ignoreNext {
			ignoreNext = false
			continue
		}
		if part == 2 {
			// check if it's end of garbage
			if inGarbage && v != '>' && v != '!' {
				characters++
			}
		}
		switch v {
		case '{':
			if !inGarbage {
				level++
			}
			break
		case '}':
			if !inGarbage {
				score += level
				level--
			}
			break
		case '<':
			inGarbage = true
			break
		case '>':
			inGarbage = false
			break
		case '!':
			ignoreNext = true
		}
	}
	if part == 1 {
		fmt.Println(score)

	} else {
		fmt.Println(characters)
	}
}
