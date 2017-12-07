package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inFile, _ := os.Open("input.txt")
	set := make(map[string]int)
	list := make([]string, 0)
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		row := strings.Fields(scanner.Text())
		if len(row) > 2 {
			val, _ := strconv.Atoi(strings.Trim(row[1], "()"))
			set[row[0]] = val
			for i := 3; i < len(row); i++ {
				// trim ',' character from words
				list = append(list, strings.Trim(row[i], ","))
			}
		}
	}
	for k, _ := range set {
		if contains(list, k) == false {
			fmt.Println(k)
			break
		}
	}
}

func contains(list []string, a string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
