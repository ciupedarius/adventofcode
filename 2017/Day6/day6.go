package main

import (
	"fmt"
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
	input := []int{11, 11, 13, 7, 0, 15, 5, 5, 4, 4, 1, 1, 7, 1, 15, 11}
	set := make(map[string]int)
	loops := 0
	for true {
		i, max := max(input)
		input = cycle(input, i, max)
		key := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(input)), ""), "[]")
		loops++
		value, ok := set[key]
		if ok {
			if part == 1 {
				fmt.Println(loops)
			} else {
				fmt.Println(loops - value)
			}
			break
		} else {
			set[key] = loops
		}
	}

}

// find first max and its position
func max(arr []int) (int, int) {
	max := arr[0]
	index := 0
	for i, n := range arr {
		if n > max {
			max = n
			index = i
		}
	}
	return index, max
}

// cycle array starting from position i with n blocks
func cycle(arr []int, i int, n int) []int {
	blocks := 0
	arr[i] = 0
	i++
	i = int(i % len(arr))
	for blocks < n {
		arr[i]++
		i++
		i = int(i % len(arr))
		blocks++
	}
	return arr
}
