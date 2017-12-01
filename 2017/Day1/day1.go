package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
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
	sum := 0
	for i, _ := range input {
		var j int
		if part == 1 {
			j = (i + 1) % len(input)
		} else {
			j = (i + len(input)/2) % len(input)
		}
		if input[i] == input[j] {
			value, _ := strconv.Atoi(string(input[i]))
			sum = sum + value
		}
	}
	fmt.Println("Sum=", sum)
}
