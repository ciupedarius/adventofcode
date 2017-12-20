package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

type coord struct {
	x, y, z int
}

type particle struct {
	p, v, a  coord
	collided bool
}

func (part particle) fromProps(point []string, velocity []string, acc []string) particle {
	p := coord{strToInt(point[0]), strToInt(point[1]), strToInt(point[2])}
	v := coord{strToInt(velocity[0]), strToInt(velocity[1]), strToInt(velocity[2])}
	a := coord{strToInt(acc[0]), strToInt(acc[1]), strToInt(acc[2])}
	return particle{p, v, a, false}
}

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
	lines := strings.Split(string(input), "\r\n")
	particles := make([]particle, 0)
	//read input
	for _, v := range lines {
		var p particle
		props := strings.Split(v, " ")
		point := strings.Split(strings.TrimSuffix(strings.TrimPrefix(props[0], "p=<"), ">,"), ",")
		velocity := strings.Split(strings.TrimSuffix(strings.TrimPrefix(props[1], "v=<"), ">,"), ",")
		acc := strings.Split(strings.TrimSuffix(strings.TrimPrefix(props[2], "a=<"), ">"), ",")
		p = p.fromProps(point, velocity, acc)
		particles = append(particles, p)
	}
	closest := 0
	// "in the long term" = 1000 iterations...
	for k := 0; k < 1000; k++ {
		min := int(^uint(0) >> 1)
		// collision set
		set := make(map[coord]int, 0)
		for i, val := range particles {
			if !val.collided {
				manhattan := math.Abs(float64(val.p.x)) + math.Abs(float64(val.p.y)) + math.Abs(float64(val.p.z))
				if min > int(manhattan) {
					min = int(manhattan)
					closest = i
				}
				velocity := coord{val.v.x + val.a.x, val.v.y + val.a.y, val.v.z + val.a.z}
				position := coord{val.p.x + velocity.x, val.p.y + velocity.y, val.p.z + velocity.z}
				particles[i] = particle{position, velocity, particles[i].a, particles[i].collided}
				// check collision
				if part == 2 {
					index, ok := set[particles[i].p]
					if !ok {
						set[particles[i].p] = i
					} else {
						// if collision happends, delete both particles
						particles[i].collided = true
						particles[index].collided = true
					}
				}
			}
		}
	}
	if part == 1 {
		fmt.Println(closest)
	} else {
		count := 0
		for _, v := range particles {
			if !v.collided {
				count++
			}
		}
		fmt.Println(count)
	}
}

func strToInt(s string) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return v
}
