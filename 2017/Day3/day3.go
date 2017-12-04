package main

import (
	"fmt"
	"math"
)

type coord struct {
	x int
	y int
}

func main() {
	input := 361527

	part1(input)
	part2(input)
}

func part1(input int) {
	// improved version
	sideLength := int(math.Ceil(math.Sqrt(float64(input))))
	if sideLength%2 == 0 {
		sideLength++ // even numbers belong to odd side lengths (3,5,7...)
	}
	y := (sideLength - 1) / 2                                              // axis length from '1'
	indexCycle := input - (sideLength-2)*(sideLength-2)                    // cycle starts when a new spiral is formed
	offsetFromCorner := indexCycle % (sideLength - 1)                      // distance to nearest corner
	x := int(math.Abs(float64(offsetFromCorner - ((sideLength - 1) / 2)))) // chage from offset to corner to offset from axis
	fmt.Println(x + y)
}

func part2(input int) {
	// compute whole spiral..
	num := 1        //current number
	level := 2      // spiral level
	count := 0      //numbers computed
	side := "right" // current side
	x := 1          //x coordinate
	y := 0          //y coordinate
	matrix := make(map[coord]int)
	matrix[coord{0, 0}] = num
	for input > num {
		// compute number and switch side when reaching corners
		switch side {
		case "right":
			num = sum_neighbor(matrix, x, y)
			matrix[coord{x, y}] = num
			count++
			if count%level == 0 {
				side = "up"
				x--
			} else {
				y++
			}
		case "up":
			num = sum_neighbor(matrix, x, y)
			matrix[coord{x, y}] = num
			count++
			if count%level == 0 {
				side = "left"
				y--
			} else {
				x--
			}
		case "left":
			num = sum_neighbor(matrix, x, y)
			matrix[coord{x, y}] = num
			count++
			if count%level == 0 {
				side = "down"
				x++
			} else {
				y--
			}
		case "down":
			num = sum_neighbor(matrix, x, y)
			matrix[coord{x, y}] = num
			count++
			// when at right corner end, begin a new level
			if count%level == 0 {
				side = "right"
				level += 2
				x++
			} else {
				x++
			}
		}
	}
	fmt.Println("Number =", num)
}

// sum of all 8 matrix neighbors
func sum_neighbor(m map[coord]int, x int, y int) int {
	s := 0
	for i := x - 1; i < x+2; i++ {
		for j := y - 1; j < y+2; j++ {
			s = m[coord{i, j}] + s
		}
	}
	return s
}
