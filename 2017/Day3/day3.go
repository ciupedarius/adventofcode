package main

import (
	"fmt"
	"math"
)

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

}
