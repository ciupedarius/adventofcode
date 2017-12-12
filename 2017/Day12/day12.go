package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var pipes = make(map[int]bool, 0)
var count int

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

	a := make([][]int, 2000)
	for i := range a {
		a[i] = make([]int, 2000)
	}

	for scanner.Scan() {
		row := strings.Fields(scanner.Text())
		x, _ := strconv.Atoi(row[0])
		for i := 2; i < len(row); i++ {
			row[i] = strings.Trim(row[i], ",")
			y, _ := strconv.Atoi(row[i])
			a[x][y] = 1
			a[y][x] = 1
		}
	}
	visited := make(map[int]bool)
	if part == 1 {
		connected := bfs(0, a, visited)
		fmt.Println(len(connected))
	} else {
		for i := 0; i < 2000; i++ {
			bfs(i, a, visited)
		}
		fmt.Println(count)
	}

}

func bfs(vert int, matrix [][]int, visited map[int]bool) map[int]bool {
	_, ok := visited[vert]
	if !ok {
		visited[vert] = true
		count++
		queue := make([]int, 0)
		queue = append(queue, vert)
		for true {
			current := queue[0]
			queue = queue[1:]
			for i := 0; i < 2000; i++ {
				if matrix[current][i] == 1 {
					_, ok := visited[i]
					if !ok {
						visited[i] = true
						queue = append(queue, i)
					}
				}
			}
			if len(queue) == 0 {
				break
			}
		}
	}
	return visited
}
