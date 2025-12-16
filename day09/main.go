package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ManuelGarciaF/AoC-2025/commons"
)

func main() {
	coords := parseInput(os.Args[1])
	fmt.Println("Part1: ", solve(coords, noFilter))
	fmt.Println("Part2: ", solve(coords, contained))
}

func solve(coords []commons.Coord, valid func(x1, x2 commons.Coord, coords []commons.Coord) bool) int {
	maximum := 0
	for i, x1 := range coords {
		for j := i + 1; j < len(coords); j++ {
			x2 := coords[j]
			area := area(x1, x2)
			// Cheapest to find out if area is bigger first
			if area > maximum && valid(x1, x2, coords) {
				maximum = area
			}
		}
	}

	return maximum
}

func noFilter(x1, x2 commons.Coord, coords []commons.Coord) bool {
	return true
}

func contained(x1, x2 commons.Coord, coords []commons.Coord) bool {
	rXMin, rXMax := min(x1.X, x2.X), max(x1.X, x2.X)
	rYMin, rYMax := min(x1.Y, x2.Y), max(x1.Y, x2.Y)

	// We must check that no edge of the polygon "cuts" the rectangle
	for i, c1 := range coords {
		c2 := coords[(i+1)%len(coords)]
		eXMin, eXMax := min(c1.X, c2.X), max(c1.X, c2.X)
		eYMin, eYMax := min(c1.Y, c2.Y), max(c1.Y, c2.Y)

		// AABB test
		if !(eXMin >= rXMax || eXMax <= rXMin || eYMin >= rYMax || eYMax <= rYMin) {
			return false
		}
	}
	return true
}

func area(a, b commons.Coord) int {
	xd := abs(a.X-b.X) + 1
	yd := abs(a.Y-b.Y) + 1
	return xd * yd
}

func abs(a int) int {
	return max(a, -a)
}

func parseInput(path string) []commons.Coord {
	f := commons.Must(os.Open(path))
	defer f.Close()

	s := bufio.NewScanner(f)

	cs := make([]commons.Coord, 0)

	for s.Scan() {
		nums := strings.Split(s.Text(), ",")
		cs = append(cs, commons.Coord{
			X: commons.MustAtoi(nums[0]),
			Y: commons.MustAtoi(nums[1]),
		})
	}

	return cs
}
