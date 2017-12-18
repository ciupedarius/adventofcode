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
	array := make([]int, 0)
	if part == 1 {
		for _, s := range strings.Split(string(input), ",") {
			val, _ := strconv.Atoi(s)
			array = append(array, val)
		}
		list := make([]int, 256)
		for i := range list {
			list[i] = i
		}
		_, _, list = hashRound(array, 0, 0, list)
		fmt.Println(list[0] * list[1])
	} else {
		part2(input)
	}
}

func hashRound(array []int, position int, skip int, list []int) (int, int, []int) {
	for _, v := range array {
		// reverse elements
		for k := 0; k < v/2; k++ {
			i := (k + position) % len(list)
			j := (v - 1 - k + position) % len(list)
			list[i], list[j] = list[j], list[i]
		}
		position = (position + v + skip) % len(list)
		skip++
	}
	return position, skip, list
}

func part2(input []byte) {
	array := make([]int, 0)
	sequence := []byte{17, 31, 73, 47, 23}
	input = append(input, sequence...)
	for _, v := range input {
		array = append(array, int(v))
	}
	position := 0
	skip := 0
	list := make([]int, 256)
	for i := 0; i < len(list); i++ {
		list[i] = i
	}
	for i := 0; i < 64; i++ {
		position, skip, list = hashRound(array, position, skip, list)
	}
	denseHash := make([]int, 16)
	for i := 0; i < 16; i++ {
		for j := i * 16; j < (i+1)*16; j++ {
			denseHash[i] = denseHash[i] ^ list[j]
		}
	}
	hash := ""
	for _, v := range denseHash {
		hex := fmt.Sprintf("%x", v)
		// add 0 padding to make a 2 digits number
		if len(hex) == 1 {
			hex = "0" + hex
		}
		hash = hash + hex
	}
	fmt.Println(hash)
}
