package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	count := 0
	for scanner.Scan() {
		row := strings.Fields(scanner.Text())
		set := make(map[string]bool)
		for i, word := range row {
			if part == 2 {
				// sort letters in word
				letters := strings.Split(word, "")
				sort.Strings(letters)
				word = strings.Join(letters, "")
			}
			// check if word is in set
			_, inSet := set[word]
			if inSet {
				break
			} else {
				set[word] = true
			}
			// if reached at end of passphrase and no breaks, inc count
			if i == len(row)-1 {
				count++
			}
		}
	}
	fmt.Println(count)
}
