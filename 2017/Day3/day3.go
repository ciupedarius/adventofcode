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
	i := 1
	sq := 1
	var x, y int
	for true {
		if input < sq {
			// find y coordinate = depth from '1'
			y = (i - 1) / 2
			// find x coordinate by comparing with axes from '1'
			//    **x**
			//	  x*1*x
			//	  **x**
			num := int(math.Sqrt(float64(sq)))
			min := sq
			for j := sq - int(num/2); j > (num-2)*(num-2); j = j - num + 1 {
				diff := int(math.Abs(float64(j - input)))
				if diff < min {
					min = diff
					x = diff
				}
			}
			break
		} else {
			i += 2
			sq = i * i
		}
	}
	fmt.Println("Distance =", x+y)
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
