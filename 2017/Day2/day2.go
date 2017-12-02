package main

import (
	"bufio"
	"fmt"
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
	inFile, _ := os.Open("input.txt")
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)
	sum := 0
	for scanner.Scan() {
		min := int(^uint(0) >> 1)
		max := 0
		value := 0
		row := strings.Fields(scanner.Text())
		for i, c := range row {
			num, _ := strconv.Atoi(c)
			if part == 1 {
				if num < min {
					min = num
				}
				if num > max {
					max = num
				}
			} else {
				for _, c := range row[i+1:] {
					num2, _ := strconv.Atoi(c)
					if num > num2 {
						if num%num2 == 0 {
							value = num / num2
							break
						}
					} else {
						if num2%num == 0 {
							value = num2 / num
							break
						}
					}

				}
			}
		}
		if part == 1 {
			sum += max - min
		} else {
			sum += value
		}
	}
	fmt.Println("Sum=", sum)
}
