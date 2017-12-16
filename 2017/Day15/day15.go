package main

import (
	"fmt"
	"os"
	"strconv"
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
	genA := 699
	genB := 124
	factorA := 16807
	factorB := 48271
	var bitsA, bitsB string
	count := 0
	pairs := 40000000
	if part == 2 {
		pairs = 5000000
	}
	for i := 0; i < pairs; i++ {
		for true {
			genA, bitsA = getNextValue(genA, factorA)
			if genA%4 == 0 || part == 1 {
				break
			}
		}
		for true {
			genB, bitsB = getNextValue(genB, factorB)
			if genB%8 == 0 || part == 1 {
				break
			}
		}

		if bitsA == bitsB {
			count++
		}
	}
	fmt.Println(count)
}

func getNextValue(current int, factor int) (int, string) {
	div := 2147483647
	next := (current * factor) % div
	nextBinary := strconv.FormatInt(int64(next), 2)
	// add 16 padding bits
	if len(nextBinary) < 16 {
		nextBinary = "00000000000000000" + nextBinary
	}
	nextBinStr := nextBinary[len(nextBinary)-16 : len(nextBinary)]
	return next, nextBinStr
}
