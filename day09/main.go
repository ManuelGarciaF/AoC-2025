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

	// Mapping to smaller numbers
	// xcoords := make([]commons.Coord, len(coords))
	// copy(xcoords, coords)
	// ycoords := make([]commons.Coord, len(coords))
	// copy(ycoords, coords)
	// slices.SortFunc(xcoords, func(a, b commons.Coord) int { return a.X - b.X })
	// slices.SortFunc(ycoords, func(a, b commons.Coord) int { return a.Y - b.Y })
	// xmap := make(map[int]int, len(xcoords))
	// for i, v := range xcoords {
	// 	xmap[v.X] = i
	// }
	// ymap := make(map[int]int, len(ycoords))
	// for i, v := range ycoords {
	// 	ymap[v.Y] = i
	// }
	// newCoords := make([]commons.Coord, len(coords))
	// for i, c := range coords {
	// 	newCoords[i] = commons.Coord{X: xmap[c.X], Y: ymap[c.Y]}
	// }

	fmt.Println("Part 1: ", part1(coords))
	fmt.Println("Part 2: ", part2(coords))
}

func part1(coords []commons.Coord) int {
	maximum := 0
	for i, x1 := range coords {
		for j := i + 1; j < len(coords); j++ {
			x2 := coords[j]
			maximum = max(maximum, area(x1, x2))
		}
	}

	return maximum
}

func part2(coords []commons.Coord) int {
	maximum := 0
	// imgcount := 0
	for i, x1 := range coords {
		for j := i + 1; j < len(coords); j++ {
			x2 := coords[j]
			area := area(x1, x2)
			// Cheapest to find out if area is bigger first
			if area <= maximum {
				continue
			}
			if !contained(x1, x2, coords) {
				// if i > 244 && imgcount < 300 {
				// 	plotBox(fmt.Sprint(i, j), x1, x2, coords)
				// 	imgcount++
				// }
				continue
			}

			fmt.Println("New best: ", x1, x2, area)
			// plotBox(fmt.Sprint("valid", i, j), x1, x2, coords)
			maximum = area
		}
	}

	return maximum
}

func contained(x1, x2 commons.Coord, coords []commons.Coord) bool {
	xMin, xMax := min(x1.X, x2.X), max(x1.X, x2.X)
	yMin, yMax := min(x1.Y, x2.Y), max(x1.Y, x2.Y)

	// We must check that no edge of the polygon "cuts" the rectangle
	for i, c1 := range coords {
		c2 := coords[(i+1)%len(coords)]

		if c1.X == c2.X {
			// Vertical edge, we check the horizontal sides of the rectangle.
			edgeX := c1.X
			if edgeX <= xMin || edgeX >= xMax {
				continue // Doesn't touch the rectangle
			}

			start := min(c1.Y, c2.Y)
			end := max(c1.Y, c2.Y)
			// Top edge (yMin)
			if start <= yMin && end > yMin {
				return false
			}
			// Bottom edge (yMax)
			if start < yMax && end >= yMax {
				return false
			}

		} else if c1.Y == c2.Y {
			// Horizontal edge, we check the vertical sides of the rectangle.
			edgeY := c1.Y
			if edgeY <= yMin || edgeY >= yMax {
				continue // Doesn't touch the rectangle
			}

			start := min(c1.X, c2.X)
			end := max(c1.X, c2.X)
			// Left edge (xMin)
			if start <= xMin && end > xMin {
				return false
			}
			// Right edge (xMax)
			if start < xMax && end >= xMax {
				return false
			}
		} else {
			panic("Invalid input")
		}
	}
	return true
}

// func insidePolygon(point commons.Coord, coords []commons.Coord) bool {
// 	inside := false

// 	px, py := point.X, point.Y
// 	for i, c1 := range coords {
// 		c2 := coords[(i+1)%len(coords)]

// 		if px == c1.X && py == c1.Y { // Is in a corner
// 			return true
// 		}

// 		// We cast a horizontal ray, must check if it collides with the edge
// 		// Ignore horizontal edges (never collides)
// 		if c1.Y == c2.Y {
// 			if c1.Y == py { // Boundary
// 				return true
// 			}
// 			continue
// 		}

// 		yMin, yMax := min(c1.Y, c2.Y), max(c1.Y, c2.Y)
// 		if yMin <= py && yMax >= py {
// 			if c1.X == px {
// 				return true // Boundary
// 			}
// 			if c1.X > px {
// 				inside = !inside
// 			}
// 		}
// 	}

// 	return inside
// }

func area(a, b commons.Coord) int {
	xd := a.X - b.X + 1
	yd := a.Y - b.Y + 1
	return max(xd*yd, -xd*yd) // Mod
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
