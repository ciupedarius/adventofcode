package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	arr := make([]int, 0)
	for _, s := range strings.Split(string(input), ",") {
		val, _ := strconv.Atoi(s)
		arr = append(arr, val)
	}
	list := make([]int, 256)
	for i := 0; i < len(list); i++ {
		list[i] = i
	}
	position := 0
	skip := 0
	for _, v := range arr {
		list = reverse(list, position, v)
		position = (position + v + skip) % len(list)
		skip++

	}
	fmt.Println(list[0] * list[1])

}

func reverse(arr []int, position int, length int) []int {
	arrLen := len(arr)
	for k := 0; k < length/2; k++ {
		i := (k + position) % arrLen
		j := (length - 1 - k + position) % arrLen
		temp := arr[i]
		arr[i] = arr[j]
		arr[j] = temp
	}
	return arr
}
