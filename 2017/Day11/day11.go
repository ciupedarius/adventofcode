package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strings"
)

/*
  \    /
   +--+
  /y   \
-+     x+-
  \z   /
   +--+
  /    \
*/
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
	directions := strings.Split(string(input), ",")
	x := 0
	y := 0
	z := 0
	max := 0
	for _, dir := range directions {
		switch dir {
		case "n":
			y += 1
			z -= 1
			break
		case "ne":
			x += 1
			z -= 1
			break
		case "se":
			x += 1
			y -= 1
			break
		case "s":
			y -= 1
			z += 1
			break
		case "sw":
			x -= 1
			z += 1
			break
		case "nw":
			x -= 1
			y += 1
			break
		}
		distance := int((math.Abs(float64(x)) + math.Abs(float64(y)) + math.Abs(float64(z))) / 2)
		if distance > max {
			max = distance
		}
	}
	if part == 1 {
		fmt.Println((math.Abs(float64(x)) + math.Abs(float64(y)) + math.Abs(float64(z))) / 2)
	} else {
		fmt.Println(max)
	}
}
