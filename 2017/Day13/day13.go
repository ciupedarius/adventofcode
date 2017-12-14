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
	firewall := make(map[int][]int)
	max := 0
	for scanner.Scan() {
		row := strings.Fields(scanner.Text())
		layer, _ := strconv.Atoi(strings.Trim(row[0], ":"))
		size, _ := strconv.Atoi(row[1])
		firewall[layer] = make([]int, 3)
		firewall[layer][0] = 0    // layer position
		firewall[layer][1] = size // layer range
		firewall[layer][2] = 0    // 0 = add, 1 = substract
		max = layer
	}
	if part == 1 {
		severity := 0
		for i := 0; i <= max; i++ {
			v, ok := firewall[i]
			if ok {
				if v[0] == 0 {
					severity = severity + i*v[1]
				}
			}
			firewall = moveOneStep(firewall)
		}
		fmt.Println(severity)
	}
	if part == 2 {
		delay := 1
		for true {
			// compute firewall with delay
			for k, v := range firewall {
				div := delay / (v[1] - 1)
				mod := delay % (v[1] - 1)
				if div%2 == 1 {
					firewall[k][0] = v[1] - 1 - mod
					firewall[k][2] = 1
				} else {
					firewall[k][0] = mod
					firewall[k][2] = 0
				}
			}
			caught := false
			for i := 0; i <= max; i++ {
				v, ok := firewall[i]
				if ok {
					if v[0] == 0 {
						caught = true
						break
					}
				}
				firewall = moveOneStep(firewall)
			}
			if !caught {
				fmt.Println(delay)
				break
			}
			delay++
		}
	}
}

func moveOneStep(firewall map[int][]int) map[int][]int {
	for k, v := range firewall {
		if v[2] == 0 {
			firewall[k][0] = v[0] + 1
			if firewall[k][0] == v[1]-1 {
				firewall[k][2] = 1
			}
		} else {
			firewall[k][0] = v[0] - 1
			if firewall[k][0] == 0 {
				firewall[k][2] = 0
			}

		}
	}
	return firewall
}
