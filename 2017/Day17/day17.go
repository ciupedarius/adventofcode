package main

import (
	"fmt"
	"os"
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
	input := 314
	arr := make([]int, 0)
	arr = append(arr, 0, 1)
	position := 1
	if part == 1 {
		for i := 2; i < 2018; i++ {
			position = ((position + input) % i) + 1
			arr = append(arr[:position], append([]int{i}, arr[position:]...)...)
		}
		fmt.Println(arr[position+1])
	} else {
		result := 0
		for i := 2; i < 50000000; i++ {
			position = ((position + input) % i) + 1
			// check if position is after 0, no need to compute array
			if position == 1 {
				result = i
			}
		}

		fmt.Println(result)
	}
}
