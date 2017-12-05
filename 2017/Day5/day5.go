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
	strList := strings.Split(string(input), "\r\n")
	list := make([]int, len(strList))
	for i, s := range strList {
		list[i], _ = strconv.Atoi(s)
	}
	i := 0
	steps := 0
	for true {
		if i >= len(list) {
			fmt.Println(steps)
			break
		}
		prev := i
		i = i + list[i]
		if part == 2 && list[prev] > 2 {
			list[prev]--
		} else {
			list[prev]++
		}
		steps++
	}
}
